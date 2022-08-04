#!/usr/bin/python
# -*- coding: utf-8 -*-

import os, sys, shutil, logging, fcntl, glob, tarfile
import yaml, json
import subprocess
import time


reload(sys)
sys.setdefaultencoding("utf-8")
logging.basicConfig(format='%(asctime)s - %(pathname)s[line:%(lineno)d] - %(levelname)s: %(message)s', level=logging.DEBUG)


def log_fatal(message):
    logging.fatal(message)
    sys.exit(1)

def log_error(message):
    logging.error(message)

def log_warn(message):
    logging.warn(message)

def log_info(message):
    logging.info(message)


# 创建或者删除目录
def oppath(path, option):
    path = path.strip()       # 除去首位空格
    path = path.rstrip("\/")  # 除去尾部"/"符号
    if not os.path.exists(path):
        if option == 'mkdir':
            os.makedirs(path)
    else:
        if option == 'remove':
            try:
                if os.path.isdir(path):
                    shutil.rmtree(path)
                elif os.path.isfile(path):
                    os.remove(path)
                elif os.path.islink(path):
                    os.remove(path)
                elif os.path.exists(path):
                    os.remove(path)
            except OSError as e:
                log_fatal(e)


# 从tar文件中提取yaml文件内容，参数为tar文件绝对路径，以及yaml文件在tar文件中的相对路径
def loadyaml_fromtar(tarFile, yamlFileName):
    with tarfile.open(tarFile, "r") as file:
        yamlContent = yaml.load(file.extractfile(yamlFileName).read(), Loader=yaml.FullLoader)
        return yamlContent


def loadyaml(yamlFile):
    with open(yamlFile, 'r') as f:
        yamlContent = yaml.load(f.read(), Loader=yaml.FullLoader)
    return yamlContent


def dumpyaml(yamlFile, yamlContent):
    with open(yamlFile, 'w+') as f:
        fcntl.flock(f, fcntl.LOCK_EX)
        yaml.dump(yamlContent, f, default_flow_style=False)
        fcntl.flock(f, fcntl.LOCK_UN)


# 将json的字符串转换成yaml格式
def loadyaml_fromjson(jsonContent):
    yamlContent = yaml.load(jsonContent, Loader=yaml.FullLoader)
    return yamlContent


# 生成json文件之前，确保目标路径未被占用
def dumpjson(jsonFile, jsonContent):
    oppath(jsonFile, 'remove')
    with open(jsonFile, 'w+') as f:
        fcntl.flock(f, fcntl.LOCK_EX)
        json.dump(jsonContent, f, encoding='utf-8')
        fcntl.flock(f, fcntl.LOCK_UN)


# 获取ec的meta.yaml文件路径，当传入version时精确匹配，未传入version时模糊匹配
def get_metaYaml_and_ecRoot(localEcRoot, shareEcRoot, shareRoot, name, version):
    metaYamlPath = ''
    ecRoot = ''
    func1 = lambda metaFilePathList: [os.path.basename(metaFilePath).replace('.meta.yaml', '').split('_')[1] for metaFilePath in metaFilePathList]
    if version:
        metaYamlFile = name + '_' + version + '.meta.yaml'
        # 如果传入version参数，先判断本地是否存在该ec
        if os.path.isfile(os.path.join(localEcRoot, 'meta', metaYamlFile)):
            metaYamlPath = os.path.join(localEcRoot, 'meta', metaYamlFile)
            ecRoot = localEcRoot
        # 如果传入version参数，再判断共享存储是否存在该ec
        elif os.path.isfile(os.path.join(shareRoot.rstrip('/'), shareEcRoot.strip('/').rstrip('/'), 'meta', metaYamlFile)):
            metaYamlPath = os.path.join(shareRoot.rstrip('/'), shareEcRoot.strip('/').rstrip('/'), 'meta', metaYamlFile)
            ecRoot = os.path.join(shareRoot, shareEcRoot.strip('/'))
        # 如果传入version参数，但本地和共享存储上都不存在该ec则报错退出
        else:
            log_fatal('name=%s version=%s 的ec并未安装，请先安装该ec' %(name, version))
    else:
        # 如果未传入version参数，先判断本地是否能够模糊匹配到该ec，并判断是否安装了多个版本的同名ec
        if len(glob.glob(localEcRoot.rstrip('/') + '/meta/' + name + '_*.meta.yaml')) > 0:
            globlist = glob.glob(localEcRoot.rstrip('/') + '/meta/' + name + '_*.meta.yaml')
            if len(globlist) == 1:
                metaYamlPath = ''.join(globlist)
                ecRoot = localEcRoot
            else:
                log_fatal('目前已经安装了%d个版本 name=%s 的ec，请指定具体的版本[%s]' %(len(globlist), name, ' '.join(func1(globlist))))
        # 如果未传入version参数，再判断共享存储是否能够模糊匹配到该ec，并判断是否安装了多个版本的同名ec
        elif len(glob.glob(shareRoot.rstrip('/') + '/' + shareEcRoot.strip('/').rstrip('/') + '/meta/' + name + '_*.meta.yaml')) > 0:
            globlist = glob.glob(shareRoot.rstrip('/') + '/' + shareEcRoot.strip('/').rstrip('/') + '/meta/' + name + '_*.meta.yaml')
            if len(globlist) == 1:
                metaYamlPath = ''.join(globlist)
                ecRoot = os.path.join(shareRoot, shareEcRoot.strip('/'))
            else:
                log_fatal('目前已经安装了%d个版本 name=%s 的ec，请指定具体的版本[%s]' %(len(globlist), name, ' '.join(func1(globlist))))
        else:
            log_fatal('name=%s 的ec并未安装，请先安装该ec' %(name))
    return metaYamlPath, ecRoot


# 执行shell命令一次，成功则返回输出，失败则退出
def exec_cmd_once(option, verbose=True):
    #status, output = commands.getstatusoutput(option)

    status = -1
    output = ''
    try:
        output = subprocess.check_output(option, shell=True)
        status = 0
    except subprocess.CalledProcessError as ex:
        output = ex.output
        status = ex.returncode

    if verbose:
        log_info("command: {}".format(option))
        log_info("exit code: {}".format(status))
        log_info("output: {}".format(output))

    return status, output


# 执行shell命令，成功则返回输出，失败则退出
def exec_cmd(option, err_tip, ok_tip='', system=False, verbose=False, ignore_error=False, retry=1):
    status = -1
    output = ''
    for i in range(1,retry+1):
        if system:
            status = os.system(option)
        else:
            status, output = exec_cmd_once(option, verbose)
        if status == 0:
            if ok_tip != '':
                log_info(ok_tip)
            return output
        else:
            if i >= retry:
                if not ignore_error:
                    log_error(err_tip)
                    sys.exit(status)
                else:
                    log_warn(err_tip)
                    return output
            else:
                log_warn(err_tip)
                time.sleep(3)


# 更新ecset中的ec信息
def update_ecinfo_in_ecset(option, value, confDir, ecName, ecVersion, cluster, keyPrefix, retry=10):
    findEc = False
    if cluster == '':
        fileList = os.listdir(confDir)
        for filename in fileList:
            fullpath = os.path.join(confDir, filename)
            if os.path.isfile(fullpath):  # 判断文件是否存在
                contentYaml = loadyaml(fullpath)
                if contentYaml.has_key('ecs'):
                    ecsInfo = []
                    needUpdate = False
                    for ec in contentYaml['ecs']:
                        if ec['name'] == ecName and ec['version'] == ecVersion:
                            findEc = True
                            if ec.has_key(option) and ec[option] == value:
                                break
                            else:
                                ec.update({option: value})
                                needUpdate = True
                        ecsInfo.append(ec)
                    if needUpdate:
                        contentYaml['ecs'] = ecsInfo
                        dumpyaml(fullpath, contentYaml)
            if findEc:
                break
    else:
        cmdExist = exec_cmd('which eietcdctl || true', '')
        if cmdExist == "":
            sys.exit(0)
        keysExist = exec_cmd('eietcdctl %s lock %s/lock -- eietcdctl %s get %s%s/ --prefix --keys-only'
                                 % (cluster, keyPrefix, cluster, keyPrefix, confDir),
                             'ecm|OPTION:%s|END|从etcd中获取%s%s/失败' % (option, keyPrefix, confDir),
                             retry=retry)
        if keysExist != '':  # 判断配置是否在etcd中存在
            for confKey in keysExist.split('\n'):
                confKey = confKey.strip()
                if confKey == '':
                    continue
                for i in range(3):
                    output = exec_cmd('eietcdctl %s lock %s/lock -- eietcdctl %s get %s --print-value-only'
                                        % (cluster, keyPrefix, cluster, confKey),
                                    'ecm|OPTION:%s|END|从etcd中获取%s失败' % (option, confKey),
                                    retry=retry)
                    contentYaml = loadyaml_fromjson(output)
                    if contentYaml and contentYaml.has_key('ecs'):
                        break
                if contentYaml and contentYaml.has_key('ecs'):
                    for ec in contentYaml['ecs']:
                        if ec['name'] == ecName and ec['version'] == ecVersion:
                            findEc = True
                            if not ec.has_key(option) or ec[option] != value:
                                exec_cmd('eietcdctl %s lock %s/lock -- ecm update-etcd --key %s --name %s --version %s --option %s --value %s'
                                             % (cluster, keyPrefix, confKey, ecName, ecVersion, option, str(value).lower()),
                                         'ecm|OPTION:%s|END|在etcd中保存新的ec配置（%s）失败' % (option, confKey),
                                         system=True, retry=retry)
                            break
                    if findEc:
                        break

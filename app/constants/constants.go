package constants

// directory and file configuration
const DirConfig = "config"
const DirProjectScripts = "scripts"
const DirProjectScriptsInstallPrefix = "i_"
const DirProjectScriptsUninstallPrefix = "d_"
const DirHelmJenkinsMaster = "/charts/jenkins-master"
const FilenameConfiguration = "k8s_jcasc_mgmt.cnf"
const FilenameConfigurationCustom = "k8s_jcasc_custom.cnf"
const FilenameJenkinsConfigurationAsCode = "jcasc_config.yaml"
const FilenameJenkinsHelmValues = "jenkins_helm_values.yaml"
const FilenameNginxIngressControllerHelmValues = "nginx_ingress_helm_values.yaml"
const FilenamePvcClaim = "pvc_claim.yaml"
const SecretsFileEncodedEnding = ".gpg"
const ScriptsFileEnding = ".sh"

// commands
const CommandMenu = "menu"
const CommandInstall = "install"
const CommandUninstall = "uninstall"
const CommandUpgrade = "upgrade"
const CommandEncryptSecrets = "encryptSecrets"
const CommandDecryptSecrets = "decryptSecrets"
const CommandApplySecrets = "applySecrets"
const CommandApplySecretsToAll = "applySecretsToAll"
const CommandCreateProject = "createProject"
const CommandCreateDeploymentOnlyProject = "createDeploymentOnlyProject"
const CommandCreateJenkinsUserPassword = "createJenkinsUserPassword"
const CommandQuit = "quit"

// helm commands
const HelmCommandInstall = "install"
const HelmCommandUpgrade = "upgrade"

// error
const ErrorPromptFailed = "prompt failed"

// colors
const ColorNormal = "\033[0m"
const ColorInfo = "\033[1;34m"
const ColorNotice = "\033[1;36m"
const ColorWarning = "\033[1;33m"
const ColorError = "\033[1;31m"
const ColorDebug = "\033[0;36m"

// kubectl field names
const KubectlOutputFieldNamespace = "NAME"
const KubectlOutputFieldPvcName = "NAME"
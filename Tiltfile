# Load ko extension for fast Go builds
load('ext://ko', 'ko_build')

# Add .tools to PATH so ko extension can find local ko binary
tools_dir = os.path.abspath('.tools')
if os.path.exists(tools_dir):
    current_path = os.getenv('PATH', '')
    os.putenv('PATH', tools_dir + ':' + current_path)

# Set KO_DOCKER_REPO to use KinD's local registry
os.putenv('KO_DOCKER_REPO', 'kind.local')

# Build controller image using ko
# ref is the image name (without ko:// prefix)
# import_path is the Go package to build
ko_build(
    'interview-reconciler-controller',
    'github.com/upbound/interview-reconciler/cmd/controller',
    deps=['cmd/', 'internal/', 'api/', 'go.mod', 'go.sum'],
)

# Install CRDs
k8s_yaml('config/crd/example.com_postgresconnections.yaml')

# Install RBAC
k8s_yaml([
    'config/rbac/service_account.yaml',
    'config/rbac/role.yaml',
    'config/rbac/role_binding.yaml',
])

# Deploy controller
k8s_yaml('config/controller/deployment.yaml')

# Configure resources
k8s_resource(
    'interview-reconciler-controller',
    port_forwards=['8080:8080', '8081:8081'],
    labels=['controller'],
)

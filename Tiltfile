# Use Kustomize
k8s_yaml(kustomize('k8s/overlays/dev'))

# Build Docker image
docker_build(
    'johnwesonga/go-http-server',
    '.',
    live_update=[
        sync('.', '/app'),
        run('go build -o server', trigger=['*.go']),
    ],
)

# Tell Tilt which resource to track
k8s_resource(
    workload='go-http-server',
    port_forwards=8080,
)
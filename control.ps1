Param(
    [string]
    $cmd = $(
        Write-Host "Error: $([char]27)[1;31mInvalid argument$([char]27)[1;m"
        [Environment]::Exit(0)
    )
)


if ($cmd -eq "destroy") {
    kubectl delete svc microservice-backend
    kubectl delete deployments microservice-backend
    kubectl delete svc microservice-frontend
    kubectl delete deployments microservice-frontend
} elseif ($cmd -eq "apply") {
    Set-Location .\src\backend
    docker build . -t microservice-backend --no-cache
    Set-Location ..\frontend\
    docker build . -t microservice-frontend --no-cache
    Set-Location ../..
    kubectl apply -f .\pods\frontend.yml
    kubectl apply -f .\pods\backend.yml
} else {
    Write-Host "Error: $([char]27)[1;31mInvalid command$([char]27)."
}

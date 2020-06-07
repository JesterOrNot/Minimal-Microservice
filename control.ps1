Param(
    [string]
    $cmd = $(
        Write-Host "Error: $([char]27)[1;31mInvalid argument$([char]27)[1;m"
        [Environment]::Exit(0)
    )
)


if ($cmd -eq "destroy") {
    bash -c "scripts/apply.sh delete"
} elseif ($cmd -eq "japply") {
    bash -c "scripts/apply.sh delete"
    bash -c "scripts/apply.sh apply" 
} elseif ($cmd -eq "apply") {
    Set-Location .\src\backend
    docker build . -t microservice-backend --no-cache
    Set-Location ..\frontend\
    docker build . -t microservice-frontend --no-cache
    Set-Location ../..
    bash -c "scripts/apply.sh delete"
    bash -c "scripts/apply.sh apply"
} else {
    Write-Host "Error: $([char]27)[1;31mInvalid command$([char]27)."
}

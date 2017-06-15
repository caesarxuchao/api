set -o errexit
set -o nounset
set -o pipefail

 # TODO: use the real one
 kubernetes_remote="https://github.com/caesarxuchao/kubernetes"
 
 git remote add upstream-kube "${kubernetes_remote}" || true
 git fetch upstream-kube
 git checkout master
 git branch -D kube-sync || true
 git checkout upstream-kube/master -b kube-sync
 
 #git filter-branch --index-filter 'git rm --cached -qr --ignore-unmatch -- . && git reset -q $GIT_COMMIT -- pkg/api pkg/apis' --prune-empty HEAD
 git filter-branch -f --index-filter 'git rm --cached -qr --ignore-unmatch -- . && git reset -q $GIT_COMMIT -- pkg/api pkg/apis' --msg-filter 'awk 1 && echo && echo "Kubernetes-commit: ${GIT_COMMIT}"' --prune-empty HEAD
 
 mkdir -p core/v1
 mv pkg/api/v1/* core/v1/

dirs=$(find pkg/apis/ -maxdepth 2 -mindepth 2 -type d -name v*)

for dir in $dirs; do
    no_pkg="${dir#pkg/apis/}"
    mkdir -p $no_pkg && mv $dir/* $no_pkg
done

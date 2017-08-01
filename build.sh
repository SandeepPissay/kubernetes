BUILD_TAG=$1
sudo make quick-release && \
cd cluster/images/hyperkube && \
sudo make build VERSION=$BUILD_TAG && \
sudo docker tag gcr.io/google_containers/hyperkube-amd64:$BUILD_TAG sandeeppissay/hyperkube-amd64:$BUILD_TAG && \
sudo docker push sandeeppissay/hyperkube-amd64:$BUILD_TAG && \
cd - && \
sudo docker run -it --rm --env="PS1=[container]:\w> " --net=host cnastorage/kubernetes-anywhere:latest /bin/bash

FROM golang:alpine AS build-env
# The GOPATH in the image is /go.
ADD . /go/src/github.com/kubeflow/katib
WORKDIR /go/src/github.com/kubeflow/katib/cmd/manager-rest
RUN if [ "$(uname -m)" = "ppc64le" ]; then \
        apk --update add gcc musl-dev  && \
        go build -o vizier-manager-rest ./v1alpha1; \
    else \
        go build -o vizier-manager-rest ./v1alpha1; \
    fi

FROM alpine:3.7
WORKDIR /app
COPY --from=build-env /go/src/github.com/kubeflow/katib/cmd/manager-rest/vizier-manager-rest /app/
ENTRYPOINT ["./vizier-manager-rest"]

FROM python:3

ADD . /usr/src/app/github.com/kubeflow/katib
WORKDIR /usr/src/app/github.com/kubeflow/katib/cmd/suggestion/bayesianoptimization/v1alpha2
RUN if [ "$(uname -m)" = "ppc64le" ]; then \
       apt-get update && apt-get -y install libblas-dev liblapack-dev libatlas-base-dev gfortran && \
       pip install cython; \
    fi
RUN pip install --no-cache-dir -r requirements.txt
ENV PYTHONPATH /usr/src/app/github.com/kubeflow/katib:/usr/src/app/github.com/kubeflow/katib/pkg/api/v1alpha2/python

ENTRYPOINT ["python", "main.py"]

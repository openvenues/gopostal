FROM golang

RUN apt-get update \
    && apt-get install --no-install-recommends -y curl autoconf automake libtool pkg-config \
    && rm -rf /var/lib/apt/lists/* \
    && apt-get purge --auto-remove -y \
    && rm -rf /src/*.deb

WORKDIR /usr/src/libpostal

RUN git clone https://github.com/openvenues/libpostal . \
    && git checkout tags/v1.1-alpha -b v1.1-alpha \
    && ./bootstrap.sh \
    && ./configure \
    && make -j4 \
    && make install

WORKDIR /usr/src/app

COPY . .

RUN go install -mod=vendor -v ./...

RUN ldconfig

ENTRYPOINT [ "postal" ]
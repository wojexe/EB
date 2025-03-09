# ubuntu:24.02 jest niedostępne na Docker Hub
FROM ubuntu:24.04

# Python
RUN apt-get update && \
    apt-get install -y software-properties-common && \
    add-apt-repository ppa:deadsnakes/ppa && \
    apt-get update && \
    apt-get install -y python3.10

# SDKMAN (Java, Kotlin, Gradle)

RUN apt-get update && \
    apt-get install -y curl zip unzip ca-certificates bash

RUN curl -s "https://get.sdkman.io" | bash

ENV SDKMAN_DIR="/root/.sdkman"
ENV PATH=${SDKMAN_DIR}/candidates/java/current/bin:${SDKMAN_DIR}/candidates/kotlin/current/bin:${SDKMAN_DIR}/candidates/gradle/current/bin:${PATH}

RUN bash -c "source ${SDKMAN_DIR}/bin/sdkman-init.sh && \
    sdk install java 8.0.392-zulu && \
    sdk default java 8.0.392-zulu && \
    sdk install kotlin && \
    sdk install gradle"

COPY build.gradle /app/
COPY src /app/src/

WORKDIR /app/

RUN bash -c "source ${SDKMAN_DIR}/bin/sdkman-init.sh && gradle build"

CMD [ \
    "bash", "-c", \
    "python3.10 --version && echo && java -version && echo && kotlin -version && echo && gradle --version && echo && source ${SDKMAN_DIR}/bin/sdkman-init.sh && gradle run" \
    ]

WORKDIR /app

FROM selenium/standalone-firefox:91.0-20210823

ARG JAR_FILE=target.jar

COPY ${JAR_FILE} app.jar


# entrypoint

ENTRYPOINT ["java","-jar","/app.jar"]
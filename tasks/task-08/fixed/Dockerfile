FROM vulhub/spring-webmvc:5.3.17

# Keep the application, JDK and Tomcat identical; update all Spring modules.
RUN set -eu; \
    mkdir -p /usr/local/tomcat/webapps/ROOT; \
    unzip -q /usr/local/tomcat/webapps/ROOT.war -d /usr/local/tomcat/webapps/ROOT; \
    cd /usr/local/tomcat/webapps/ROOT/WEB-INF/lib; \
    for module in aop beans context core expression jcl web webmvc; do \
      curl --fail --silent --show-error --location --retry 3 \
        "https://repo.maven.apache.org/maven2/org/springframework/spring-${module}/5.3.18/spring-${module}-5.3.18.jar" \
        -o "spring-${module}-5.3.18.jar"; \
      rm "spring-${module}-5.3.17.jar"; \
    done; \
    rm -f /usr/local/tomcat/webapps/ROOT.war

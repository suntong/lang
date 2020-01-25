CLASSFDIR=/java/classes
JAVAFILES= $(wildcard *.java)
CLASSFILES= $(subst .java,.class,$(JAVAFILES))

all: ${CLASSFILES}

${CLASSFDIR}/%.class: %.java
	javac -d ${CLASSFDIR} $<

%.class: %.java
	javac -d ${CLASSFDIR} $<

clean:
	rm *~ ${CLASSFILES}

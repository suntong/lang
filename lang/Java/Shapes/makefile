.SUFFIXES: .java .class

.java.class:
	javac $<

all:	Shape.class Shapes.class

test:	all
	@echo " == Testing == "
	java Shapes

clean:
	rm *.class *~


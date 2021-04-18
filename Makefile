PANDOC=pandoc -s --webtex -i -t slidy

html: doc/slides-de.md generated cp-pics
	$(PANDOC) doc/slides-de.md -o generated/slides-de.html
	cp doc/flowchart.svg generated/

cp-pics: doc/flowchart.svg
	cp doc/flowchart.svg generated/

generated:
	mkdir generated
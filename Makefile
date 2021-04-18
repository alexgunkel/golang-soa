PANDOC=pandoc -s --webtex -i -t slidy

html: doc/slides-de.md generated
	$(PANDOC) doc/slides-de.md -o generated/slides-de.html

generated:
	mkdir generated
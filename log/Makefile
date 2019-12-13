# Copyright © 2019 Ruben Mosblech
# This file belongs to the creator
# You may use, modify and distribute this file as long as you keep this notice

MAKEFILENAME	= go.make
MAKEFILE-GO		= $(MAKEFILENAME)
MAKEFLAGS		+= --no-print-directory

.PHONY: all
all: $(MAKEFILE-GO)
	@$(MAKE) -f $(MAKEFILENAME) all || ret=$$?; exit $$ret

%: | $(MAKEFILE-GO)
	@$(MAKE) -f $(MAKEFILENAME) "$@" || ret=$$?; exit $$ret

_clean: $(MAKEFILE-GO)
	$(MAKE) -f $(MAKEFILENAME) clean || ret=$$?; exit $$ret

clean: _clean ## run target clean in go.make and delete go.make afterwards
	rm -vf $(MAKEFILENAME)

$(MAKEFILE-GO):
	@curl -o $(MAKEFILENAME) https://gist.githubusercontent.com/h8ckr/27261b540607de4e366f59c508993ea5/raw/092b7acd2c5d4b2959cc29fa6c81dcaf4c577a7c/go.make &>/dev/null

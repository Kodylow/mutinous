cln:
  lightningd --lightning-dir=.lightning/ --signet --log-level=debug --disable-plugin bcli
alias:
  alias lc="lightning-cli --lightning-dir=.lightning/ --signet"
cln-main:
  lightningd --lightning-dir=.lightning/ --log-level=debug --disable-plugin bcli
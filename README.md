# Mutinous: Mutinynet Lightning Addresses for Everyone!

Mutinous provides mutinynet lightning addresses based off your Replit username. 

```
replitusername@mutinous.replit.app
```

You can try it by first hitting a GET against:

```
https://mutinous.replit.app/.well-known/lnurlp/replitusername

Returns:

{
  "status": "OK",
  "tag": "payRequest",
  "commentAllowed": 255,
  "callback": "http://localhost:8080/lnurlp/kody/callback",
  "metadata": "[[\"text/identifier\",\"kody@domain.com\"],[\"text/plain\",\"Sats for kody\"]]",
  "minSendable": 1000,
  "maxSendable": 110000,
  "payerData": {
    "name": {
      "mandatory": false
    },
    "email": {
      "mandatory": false
    },
    "pubkey": {
      "mandatory": false
    }
  },
  "nostrPubkey": "",
  "allowsNostr": false
}
```

Which will give you a callback to GET against with an amount >= 1000 msats

```
https://mutinous.replit.app/lnurlp/replitusername/callback?amount=1000

Returns:

{
  "status": "OK",
  "successAction": {
    "tag": "message",
    "message": "Walk the plank, this Mutiny's just getting started!"
  },
  "verify": "https://mutinous.replit.app/lnurlp/kody/verify/3L9aEJTF",
  "routes": null,
  "pr": "lntbs10n1pjt3eyhsp5q7dqq2m4xpukyhxhtlk2c6z97qhxkmdm5829gnq2hxgwzj27734qpp5fynuw63hcwnyn2eyrhghvcmmsqd4jzkt4u6u33jeanjsa4vv4geqdrltddjyar90p6z76tyv4h8g6txd9jhyg3vyf4k7eregpkh2arfdehh2uewwfjhqmrfwshxzursyfwjckezw3jhsap0wpkxz6twygkzy5mpw3ejqen0wgsxkmmy0y396hgcqp29qxpqysgqlfm2k2zcdhfpm3ul05cq8htvncsanrgw5uudkrhqt0edcvzd4wh5l6qnglfp2s85ug4f95tn80gdm8nguemyyk5ntlgajd29yved8ecqwwymne"
}

```

Which will give you a verify URL you can poll against client side to check if the invoice (pr) is paid:

```
https://mutinous.replit.app/lnurlp/lnurlp/replitusername/verify/{label}

Returns:
```


It works for any replit user and lets them generate mutinynet lightning invoices easily for building Fedi mods and other webln based applications.

As of right now when the node running the lightning address server just maintains a sqlite table for balances and their's no way to withdraw. Coming soon though!


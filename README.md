# Mutinous: Mutinynet Lightning Addresses for Everyone!

Mutinous provides mutinynet lightning addresses based off your Replit username. 

```
replitusername@mutinynet.app
```

You can try it by first hitting a GET against:

```
https://mutinynet.app/.well-known/lnurlp/replitusername

Returns:

{
  "status": "OK",
  "tag": "payRequest",
  "commentAllowed": 255,
  "callback": "http://mutinynet.app/lnurlp/kody/callback",
  "metadata": "[[\"text/identifier\",\"kody@mutinynet.app\"],[\"text/plain\",\"Sats for replitusername\"]]",
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
https://mutinynet.app/lnurlp/replitusername/callback?amount=1000

Returns:

{
  "status": "OK",
  "successAction": {
    "tag": "message",
    "message": "Walk the plank, this Mutiny's just getting started!"
  },
  "verify": "https://mutinynet.app/lnurlp/replitusername/verify/Cef5g34rpgI1F0voSyzRO8EafxSz8txF",
  "routes": null,
  "pr": "lntbs10n1pjva0nxsp5c7xdh8fgpx4g0lngyd80u3chhunk33q87s2qdtrat2yfp2gjpj8spp5dxafw2w6gwdnn554eurndyxdzy8w6qwcfx53p0plz7ctyh6g7tvqdr4tddjyar90p6z76tyv4h8g6txd9jhyg3vyf4k7eregpkh2arfdeukuet59eshqupzt5k9kgn5v4u8gtmsd3skjm3z9s39xct5wvsxvmmjyp4k7ereyfw46cqp2rzjqftf3ny6cc3lt5d67433puh3kcklllhy8l7dktpqej65racyjl25qpynmuqqqqgqqyqqqqlgqqqqqqgq2q9qxpqysgqc8nu0xgwa2h2vcqxpqgfura5g0km2ye0kptcfzz3r2xxmtd7lca8a9gpnu7pus84f3qkx6dv2sp4sxp4cwpkhrzm99yrqzhqvj3j0dqp4wgh6n"
}

```

Which will give you a verify URL you can poll against client side to check if the invoice (pr) is paid:

```
https://mutinynet.app/lnurlp/lnurlp/{replitusername}/verify/{label}

Returns (unpaid):

{
status: "OK",
settled: false,
preimage: null,
pr: "lntbs10n1pjva0nxsp5c7xdh8fgpx4g0lngyd80u3chhunk33q87s2qdtrat2yfp2gjpj8spp5dxafw2w6gwdnn554eurndyxdzy8w6qwcfx53p0plz7ctyh6g7tvqdr4tddjyar90p6z76tyv4h8g6txd9jhyg3vyf4k7eregpkh2arfdeukuet59eshqupzt5k9kgn5v4u8gtmsd3skjm3z9s39xct5wvsxvmmjyp4k7ereyfw46cqp2rzjqftf3ny6cc3lt5d67433puh3kcklllhy8l7dktpqej65racyjl25qpynmuqqqqgqqyqqqqlgqqqqqqgq2q9qxpqysgqc8nu0xgwa2h2vcqxpqgfura5g0km2ye0kptcfzz3r2xxmtd7lca8a9gpnu7pus84f3qkx6dv2sp4sxp4cwpkhrzm99yrqzhqvj3j0dqp4wgh6n"
}

or if paid returns:

{
status: "OK",
settled: true,
preimage: "b103b4e4c814c96f9982560391778f6d3c5fcec8c5323dd9b7162692f9d4f6bd",
pr: "lntbs10n1pjva0nxsp5c7xdh8fgpx4g0lngyd80u3chhunk33q87s2qdtrat2yfp2gjpj8spp5dxafw2w6gwdnn554eurndyxdzy8w6qwcfx53p0plz7ctyh6g7tvqdr4tddjyar90p6z76tyv4h8g6txd9jhyg3vyf4k7eregpkh2arfdeukuet59eshqupzt5k9kgn5v4u8gtmsd3skjm3z9s39xct5wvsxvmmjyp4k7ereyfw46cqp2rzjqftf3ny6cc3lt5d67433puh3kcklllhy8l7dktpqej65racyjl25qpynmuqqqqgqqyqqqqlgqqqqqqgq2q9qxpqysgqc8nu0xgwa2h2vcqxpqgfura5g0km2ye0kptcfzz3r2xxmtd7lca8a9gpnu7pus84f3qkx6dv2sp4sxp4cwpkhrzm99yrqzhqvj3j0dqp4wgh6n"
}
```


It works for any replit user and lets them generate mutinynet lightning invoices easily for building Fedi mods and other webln based applications.

As of right now when the node running the lightning address server just maintains a sqlite table for balances and their's no way to withdraw. Coming soon though!


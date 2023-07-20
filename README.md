# Mutinous: Mutinynet Lightning Addresses for Everyone!

Mutinous provides mutinynet lightning addresses based off your Replit username. 

```
replitusername@mutinous.replit.app
```

You can try it by first hitting a GET against:

```
https://mutinous.replit.app/.well-known/lnurlp/replitusername
```

Which will give you a callback to GET AGAINST with an amount >= 1000 msats

```
https://mutinous.replit.app/lnurlp/replitusername/callback?amount=1000

Returns:

```

Which will give you a verify URL you can poll against client side to check if the invoice is paid:

```
https://mutinous.replit.app/lnurlp/lnurlp/replitusername/verify/{label}

Returns:
```


It works for any replit user and lets them generate mutinynet lightning invoices easily for building Fedi mods and other webln based applications.

As of right now when the node running the lightning address server just maintains a sqlite table for balances and their's no way to withdraw. Coming soon though!


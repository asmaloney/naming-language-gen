[![GitHub](https://img.shields.io/github/license/asmaloney/naming-language-gen)](LICENSE) ![Build](https://github.com/asmaloney/naming-language-gen/actions/workflows/build.yaml/badge.svg)

# Naming Language Gen

This is a quick port of [Martin O'Leary](https://mewo2.com)'s [naming-language](https://github.com/mewo2/naming-language) code from JavaScript to Go. He describes in detail what it does in his [blog post](https://mewo2.com/notes/naming-language/).

**Why?**

Why not? I felt like it.

**What does it do?**

It generates a naming language based on a bunch of tables which specify phonemes & orthographic mappings along with some other rules.

**What's a naming language?**

This is a language which is intended to look good on paper, but doesn't have any real semantics (meaning) behind it. It might be used to name cities, regions, or land features on a fantasy map for example. Because the words are all generated from the same rules, they should be consistent with each other.

**Should I use this code to learn Go?**

Nope. Absolutely not. This is pretty non-idiomatic Go code. It might even be anti-idiomatic.

I started with a direct port of the JavaScript, then moved some things around and added a couple of capabilities like choosing random syllable structures from a list.

**What does it looks like?**

Building it and running it generates a random language based on internal tables and spits out some words and names like this:

```
$ go build && ./naming-language-gen
[words]: kun, eeak, iwt, mahkuz, tužnaskun, hjužkuziwt, koemahhjuž, hemwepsmot
[names]: Tounak, Tužnaskun, Hemwepsmot, Paoshaš, Hwotkun, Eeak, a Kunkiw, Itmah,
Kun Iwt, Hjužkun, Kuzužak, Eki-Kunkuz, Mahkuz, a-Kunkuz, a Iwt, Ipmonpuwn,
Mjemiwt, Iwtmiu, a Kunkuz, Umsmot
-> apply ortho: true
-> apply morph: true
-> phonemes:
     C:  ptkmnh
     V:  aeiou
     S:  sʃ
     F:  sʃzʒ
     L:  wj
-> restrictions:
-> consonant ortho:
     j  =>  j
     ʃ  =>  š
     ʒ  =>  ž
     ʧ  =>  č
     ʤ  =>  ǧ
-> vowel ortho:
     A  =>  au
     E  =>  ei
     I  =>  ie
     O  =>  ou
     U  =>  oo
-> morphemes:
>  'the'
     a
>  'words'
     kun, ak, iwt, mah, smot, shaš, kok, te, ip
>  ''
     e, kuz, tuž, nas, hjuž, koe, he, mwep, tis, pajp, mwez, hwot, mjem, mon,
     toun, tuz, pao, iwh, pwek, už, šnes, tus, nit, kiw, kez, ki, it, miu, oh,
     en, op, puwn, nwih, han, ke, um, ket, hwak, mek, pož
>  'of'
     up
```

**Huh. Weird. Can you give me another example?**

Sure.

```
$ go build && ./naming-language-gen
[words]: igtad, nartad, sumskan, tadsurk, skan, tadirm, urdtad, schiskan,
tadig, tad
[names]: Tadig, Schiskan, Skan Tad, Igtad, Tadirm, Unulmingstu, mu Stugam,
Surkumulstu, Nartad, Urdtad, Unstu, Skanig, Tadsurk, mu Tad, Imkiasguag,
Sumskan, Tad Saschan, Tadkim, Artstu, Tad-Saschan
-> apply ortho: true
-> apply morph: true
-> phonemes:
      C:  tkdgmnsʃ
      V:  aiu
      S:  sʃ
      F:  mnŋ
      L:  rl
-> restrictions: [sʃf][sʃ], [rl][rl]
-> consonant ortho:
      x  =>  ch
      ʃ  =>  sch
      ʒ  =>  zh
      ʧ  =>  tsch
      ʤ  =>  dz
      j  =>  j
-> vowel ortho:
      I  =>  ii
      O  =>  oo
      U  =>  uu
      A  =>  aa
      E  =>  ee
-> morphemes:
 >  'words'
      tad, skan, stu, gut, guag
 >  'of'
      sasch
 >  'the'
      mu
 >  ''
      ig, nar, sum, surk, irm, urd, schi, gan, nusch, sar, ti, kasch, ming,
      kas, nug, un, ul, gam, isch, giis, um, art, schan, im, kim, kias, sa,
      kang, saim, schim, sching, tik, urm, dik
```

**That's it?**

Yep.

**Why doesn't the program take any input?**

That would be the next logical step. The internal tables should be made external so they can be easily changes and fed into the program along with a language description. This would allow experimenting, fine-tuning, and reusing of language descriptions that work well.

**What else could it do?**

Martin has some suggestions at the bottom of his [post](http://mewo2.com/notes/naming-language/) which outline other things that could be added.

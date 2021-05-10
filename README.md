# universalang

A silly attempt to do an universal language, which would use for each word, an "average" of the letters/sounds of the same word in many languages around the world.

I'm not a linguist, so it's probably full of approximations, and it's still in early development.

### How to run

To build:
```
$ go build -o universalang ./src
```

Examples:
```
$./universalang "I love eating salads"
je aibii take~ sala$

$ ./universalang "Hello world, i am a translation into an universal language, this is funny. Do you understand me?"
hal muske, je ei translatio eno universal langa, ke ei fushje. do tu foeem me?

$ ./universalang "The African fish eagle is a large species of eagle found throughout sub-Saharan Africa wherever large bodies of open water with an abundant food supply occur."
afrikan pib oao ei band espe de oao faka* tartut sub-saharien africa doke grand bodi$ de oaert ma moth abunda mai kaea oiuo.

$ ./universalang "St Cyprian's Church is an Anglican parish church in the Marylebone district of London. Designed by Sir Ninian Comper in a Perpendicular Gothic style, the building was constructed between 1901 and 1903."
st cyprian ed kglrksja ei anglican paruk kerke da marylebone distriw de london. desin* bko sunshar ninian comper da perpendicu gotic ginkshe, kete ei* konstruin* beea 1901 aa 1903.
```

Because of the way the algorithm currently works, each execution can return different values. Use -v for verbose mode.

### Language characteristics

(May evolve a lot at this point)

 * All characters use ascii.
 * Plural form is represented with the suffix '$'
 * Past tense is represented with the suffix '*'.
 * Gerund is represented with the suffix '~'.
 * Articles "A"/"the" do not exist. 
 * Punctuation is basically the same than in english.
 * There is no uppercase.

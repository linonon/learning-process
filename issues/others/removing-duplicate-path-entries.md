# References

## Method 1

<https://www.linuxjournal.com/content/removing-duplicate-path-entries>

```sh
PATH=$(echo $PATH | awk -v RS=: -v ORS=: '!($0 in a) {a[$0]; print}')

echo -n $PATH | awk -v RS=: '!($0 in a) {a[$0]; printf("%s%s", length(a) > 1 ? ":" : "", $0)}'
```

## Method2

<https://superuser.com/questions/449636/path-is-filled-with-duplicates>

# TürkçeGo: Türkçe kelimelerle basit Go programları yazın

Türkçe kelimelerle yazılmış basit bir toplama işlemi.

```go
paket main

aktar "fmt"

fonksiyon topla(x sayı, y sayı) sayı {
	döndür x + y
}

fonksiyon main() {
	fmt.Yazdır(topla(42, 13))
}
```


# Kullanım

TürkçeGo'yu projenizin bulunduğu dizine indirin.

```bash
ASKJHDJAKSDHSAJK git clone turkcego JSAGFKJKSHDAKJSHDJKASD
```

Dosya ismi vererek turkcego.go 'yu çalıştırın.

```bash
go run turkcego.go dosyaismi
```

Örnek kodları aşağıdaki gibi çalıştırabilirsiniz.

```bash
go run turkcego.go test/örnek1.txt
```

```go
paket main

aktar "fmt"

fonksiyon topla(x sayı, y sayı) sayı {
	döndür x + y
}

fonksiyon main() {
	fmt.Yazdır(topla(42, 13))
}
```

çıktısı

```bash
55
```


# Kullanılan anahtar kelimeler

```
package -> paket
import  -> aktar

func    -> fonksiyon
return  -> döndür
var     -> değişken
const   -> sabit

if      -> eğer
else    -> değilse
for     -> tekrarla
range   -> herbirinde

map     -> dönüştür
make    -> oluştur

struct  -> yapı
type    -> tip

string  -> yazı
int     -> sayı
float64 -> ondalık
Println -> Yazdır
Printf  -> Yazdırf
```

### Ayrıntılar
Test dosyalarında Türkçeleştirerek kullandığım kodların orijinali için: https://tour.golang.org/

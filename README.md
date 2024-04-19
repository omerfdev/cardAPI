# Kart Türü Kontrolü API

Bu basit API, kullanıcıların bir kredi kartı numarasının ilk rakamına göre kart türünü belirler. Kart numarasının ilk rakamı incelenerek Visa veya Mastercard olup olmadığı kontrol edilir.

## Kullanım

- API, `/checkCardType` endpoint'i üzerinden kart türü kontrolü yapar.
- Kart numarası, URL parametresi olarak `cardNumber` olarak belirtilir.
- Örneğin, `http://localhost:8080/checkCardType?cardNumber=1234567890123456` adresine istek göndererek kart türünü kontrol edebilirsiniz.
- API, kart numarasının ilk rakamına göre Visa, Mastercard veya bulunamadı ("Not found") olarak yanıt verir.

## Dikkat

- Kart numarası tam olarak 16 haneli olmalıdır.
- Kart numarasının ilk rakamı incelenir ve buna göre kart türü belirlenir.
- Hatalı isteklerde uygun hata mesajları döner.

## Lisans

Bu proje MIT lisansı altında lisanslanmıştır. Daha fazla bilgi için `LICENSE` dosyasını inceleyebilirsiniz.

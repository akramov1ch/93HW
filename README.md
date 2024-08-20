# Uy vazifasi: Rate Limiting va CORS Preflight Requests bilan Go RESTFul API yaratish

## Maqsad
Ushbu uyga vazifa sizga `Go` tilida `RESTFul API` yaratish va uni `CORS` preflight so'rovlarini qo'llab-quvvatlash bilan birga, `Rate Limiting` bilan himoya qilishni qollashingiz kerak. Foydalanuvchilarning so'rovlarini monitoring qilish va cheklash uchun `Redis`-dan foydalaniladi.

## Talablar
1. **RESTFul API: Go yordamida RESTFul API yaratish:**
2. **CORS Preflight So'rovlari**:
    - `API`-ga yuboriladigan `OPTIONS` so'rovlarini aniqlang va ularga javob qaytaring.
    - Dinamik ravishda turli domenlardan keladigan so'rovlar uchun `Access-Control-Allow-Origin` headerlarini sozlang.
    
3. **Rate Limiting Middleware:**:
    - Har bir foydalanuvchi uchun so'rovlar sonini cheklang.
    - `Redis`-dan foydalanib, so'rovlar cheklovini saqlang va boshqaring.    

4. **User API:**:
    - Foydalanuvchilar uchun autentifikatsiya va autorizatsiya jarayonlarini qo'shing. Foydalanuvchilarga ma'lum so'rovlar asosida ruxsat bering. 

**Bonus Topshiriqlar**:
    - Dinamik tarzda foydalanuvchi so'rovlarini boshqarish va `Redis` orqali belgilangan vaqt davomida Rate Limiting o'rnatish.
    - Preflight so'rovlarida maxsus `Access-Control-Allow-Methods` va `Access-Control-Allow-Headers` boshqaruvini sozlash.

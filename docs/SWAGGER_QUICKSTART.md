# 🚀 Swagger UI Quick Start

## Access Swagger UI

Once your server is running, visit:

```
http://localhost:8080/swagger
```

## 3-Step Quick Test

### 1️⃣ Register & Get Token

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

Copy the `token` from the response.

### 2️⃣ Authorize in Swagger UI

1. Click the **"Authorize"** button (🔓 green lock icon)
2. Enter: `Bearer YOUR_TOKEN_HERE`
3. Click **"Authorize"** then **"Close"**

### 3️⃣ Test an Endpoint

1. Click `GET /balance`
2. Click **"Try it out"**
3. Click **"Execute"**
4. See your balance! ✅

## What You Can Do

| Feature | Endpoint | Description |
|---------|----------|-------------|
| 💰 Check Balance | `GET /balance` | View current account balance |
| 💸 Add Transaction | `POST /transactions` | Record income or expense |
| 📁 Manage Categories | `GET /categories` | View spending categories |
| 🛒 Track Big Buys | `POST /big-buys` | Record large purchases |
| ⚙️ Update Settings | `PATCH /account/timezone` | Change timezone |

## Example: Create a Transaction

1. Get a category ID:
   - `GET /categories` → Copy any category `ID`

2. Create transaction:
   ```json
   {
     "category_id": "paste-category-id-here",
     "amount": -500,
     "date": "2024-01-15T10:30:00Z",
     "note": "Coffee"
   }
   ```

3. Check balance:
   - `GET /balance` → Should show -500

## Tips

💡 **Negative amounts** = expenses  
💡 **Positive amounts** = income  
💡 **Tokens expire** after 24 hours  
💡 **All times** are in UTC  

## Need Help?

- 📖 Full docs: [docs/README.md](./README.md)
- 📋 Setup guide: [SWAGGER_SETUP.md](../SWAGGER_SETUP.md)
- 🐛 Issues? Check `docker-compose logs api`

---

**Happy Testing! 🎉**

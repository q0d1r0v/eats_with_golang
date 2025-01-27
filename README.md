# eats_with_golang

# API Routes

This document lists all the API routes for the **Eats Backend API**.

---

## Authentication Routes

| Method | Endpoint              | Description          |
| ------ | --------------------- | -------------------- |
| POST   | `/auth/register`      | Register a new user. |
| POST   | `/auth/login`         | Log in as a user.    |
| POST   | `/auth/courier/login` | Log in as a courier. |

---

## Admin Routes

### Branch Management

| Method | Endpoint                      | Description            |
| ------ | ----------------------------- | ---------------------- |
| POST   | `/admin/api/v1/create/branch` | Create a new branch.   |
| GET    | `/admin/api/v1/load/branches` | Get all branches.      |
| PUT    | `/admin/api/v1/update/branch` | Update branch details. |
| DELETE | `/admin/api/v1/delete/branch` | Delete a branch.       |

### Category Management

| Method | Endpoint                        | Description              |
| ------ | ------------------------------- | ------------------------ |
| GET    | `/admin/api/v1/load/categories` | Get all categories.      |
| POST   | `/admin/api/v1/create/category` | Create a new category.   |
| PUT    | `/admin/api/v1/update/category` | Update category details. |
| DELETE | `/admin/api/v1/delete/category` | Delete a category.       |

### Food Management

| Method | Endpoint                    | Description                 |
| ------ | --------------------------- | --------------------------- |
| GET    | `/admin/api/v1/load/foods`  | Get all foods (admin view). |
| POST   | `/admin/api/v1/create/food` | Add a new food item.        |
| PUT    | `/admin/api/v1/update/food` | Update food details.        |
| DELETE | `/admin/api/v1/delete/food` | Delete a food item.         |

### Role Management

| Method | Endpoint                    | Description          |
| ------ | --------------------------- | -------------------- |
| GET    | `/admin/api/v1/load/roles`  | Get all roles.       |
| POST   | `/admin/api/v1/create/role` | Create a new role.   |
| PUT    | `/admin/api/v1/update/role` | Update role details. |
| DELETE | `/admin/api/v1/delete/role` | Delete a role.       |

### Courier Management

| Method | Endpoint                       | Description                  |
| ------ | ------------------------------ | ---------------------------- |
| GET    | `/admin/api/v1/load/couriers`  | Get all couriers.            |
| POST   | `/admin/api/v1/create/courier` | Create a new courier.        |
| PUT    | `/admin/api/v1/update/courier` | Update courier details.      |
| DELETE | `/admin/api/v1/delete/courier` | Delete a courier (optional). |

---

## Courier Routes

| Method | Endpoint                          | Description              |
| ------ | --------------------------------- | ------------------------ |
| GET    | `/courier/api/v1/load/all/orders` | Get all assigned orders. |
| PUT    | `/courier/api/v1/update/order`    | Update order status.     |

---

## API User Routes

### Food Routes

| Method | Endpoint             | Description              |
| ------ | -------------------- | ------------------------ |
| GET    | `/api/v1/load/foods` | Get all available foods. |
| POST   | `/api/v1/order/food` | Place a food order.      |

---

## Middleware Used

- **JWTAdminMiddleware**: Applied to `/admin/` routes.
- **JWTCourierMiddleware**: Applied to `/courier/` routes.
- **JWTAuthMiddleware**: Applied to `/api/` routes.

---

## Notes

- **Base URL**: `http://localhost:8000`
- Ensure a valid JWT token is passed in the `Authorization` header for protected routes:

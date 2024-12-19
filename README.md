# todo-list-fullstack [practicode-3](https://learn.malkabruk.co.il/practicode/projects/pract-3/)

[frontend](https://todo-list-fullstack-frontend.onrender.com)

[backend](https://todo-list-fullstack-backend.onrender.com)

[database](https://console.clever-cloud.com/users/me/addons/addon_9361bbf6-4a08-4be2-8733-9e42cccbf62c)

### Project Overview

The project is a **ToDo List Fullstack application** built with a Vue.js frontend and a Go-Gin backend, utilizing MySQL as the database.

---

### Bullet Points

1. **Database Connection with GORM**:

   - Used GORM as the ORM for MySQL to simplify database interactions.
   - Automatically migrated models (`Task` and `User`) to keep the schema synchronized with the Go structs, unlike the manual scaffold approach in the instructions.

2. **DB-First vs. Code-First**:

   - Followed a **code-first approach** using GORM's auto-migration to generate database tables directly from the Go models (`Task` and `User`).
   - The instructions suggest a **DB-first approach**, where the schema is created manually or via MySQL Workbench, and scaffolding generates the ORM classes.

   **Practical Influence**:

   - Code-first ensures changes in models are reflected in the database automatically.
   - Backend developers don't need to manage separate SQL scripts or manually synchronize schema changes.

3. **MySQL DB in Clever Cloud**:

   - Configured the MySQL database hosted on Clever Cloud with the following specifications:

   **Practical Influence**:

   - **Shared vCPUS and Memory**: Performance might be affected if other users on the shared resource are consuming high CPU/memory.
   - **Max Connection Limit (5)**: Limits concurrent database operations, requiring developers to optimize database usage or use connection pooling.
   - **Max DB Size (10 MB)**: Constraints the amount of data the application can store, influencing data retention and archiving strategies.
   - **Daily Backups**: Provides disaster recovery but with limited flexibility (7 retained backups).

4. **Backend and Frontend Deployment in Render**:

   - Deployed the backend as a **web service**.
   - Deployed the frontend as a **static site** in Render, serving the Vue.js application efficiently with built-in CDN support.

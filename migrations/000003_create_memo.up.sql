CREATE TABLE memos (
    id BIGSERIAL PRIMARY KEY,

    title VARCHAR(255),
    content TEXT,

    is_public BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    
    user_id BIGINT NOT NULL
        REFERENCES users(id)
        ON DELETE RESTRICT
);
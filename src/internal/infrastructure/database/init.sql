CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    login TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

/*CREATE TABLE IF NOT EXISTS games (
    id UUID PRIMARY KEY,
    board JSONB NOT NULL,
    player1_id UUID REFERENCES users(id), 
    player2_id UUID REFERENCES users(id), 
    status TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);*/

CREATE TABLE IF NOT EXISTS games (
    id UUID PRIMARY KEY,
    board JSONB NOT NULL,
    status TEXT NOT NULL, 
    
    player_x_id UUID NOT NULL REFERENCES users(id),
    player_o_id UUID REFERENCES users(id), 
    
    current_turn_id UUID REFERENCES users(id), 
    winner_id UUID REFERENCES users(id), 
    
    is_vs_computer BOOLEAN DEFAULT FALSE,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    finished_at TIMESTAMP WITH TIME ZONE 
);

INSERT INTO users (id, login, password_hash, created_at)
VALUES ('00000000-0000-0000-0000-000000000000', 'bot', 'system_bot_no_password', NOW())
ON CONFLICT DO NOTHING;

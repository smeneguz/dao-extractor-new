-- Historical token price data (daily granularity).
-- Populated by external price fetcher (CoinGecko / DeFiLlama).
CREATE TABLE IF NOT EXISTS token_prices (
    id BIGSERIAL PRIMARY KEY,
    token_symbol TEXT NOT NULL,
    price_usd NUMERIC NOT NULL,
    market_cap_usd NUMERIC,
    volume_24h_usd NUMERIC,
    price_date DATE NOT NULL,
    source TEXT NOT NULL DEFAULT 'coingecko',
    fetched_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_token_price_date UNIQUE (token_symbol, price_date, source)
);
CREATE INDEX IF NOT EXISTS idx_token_prices_symbol ON token_prices(token_symbol);
CREATE INDEX IF NOT EXISTS idx_token_prices_date ON token_prices(price_date);

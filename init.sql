\c goldwatcher;

CREATE TABLE public.prices (
    id bigserial NOT NULL, last_at timestamptz NULL, ayar22_altin int8 NULL, ceyrek int8 NULL, yarim int8 NULL, tam int8 NULL, cumhuriyet int8 NULL, iab_kapanis int8 NULL, CONSTRAINT uni_prices_id PRIMARY KEY (id)
);
CREATE TABLE new_zealand (
    id SERIAL PRIMARY KEY,
    city TEXT,
    salary_month_nzd INT,
    rains_in_year INT
);
INSERT INTO new_zealand (city, salary_month_nzd, rains_in_year) VALUES ('Auckland', 9083, 130), ('Queenstown', 8241, 96), ('Christchurch', 8619, 83), ('Wellington', 8264, 128);
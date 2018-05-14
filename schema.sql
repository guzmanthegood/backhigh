DROP TABLE highlights;

CREATE TABLE highlights (
    id              integer         NOT NULL,
    created_at      timestamp       NOT NULL,
    title           varchar(255)    NOT NULL,
    country         varchar(255)    NOT NULL,
    price           decimal(5,2)            ,
    content         text            
);

INSERT INTO highlights (id, created_at, title, country, price, content) VALUES
(1, current_timestamp, 'Pashupatinath Temple', 'Nepal', 10.00, 'One of the most sacred Hindu temples of Nepal – Pashupatinath Temple is located on both banks of Bagmati River on the eastern outskirts of Kathmandu.'),
(2, current_timestamp, 'Swayambhunath Temple', 'Nepal', 1.50, 'Swayambhunath is an ancient religious architecture atop a hill in the Kathmandu Valley, west of Kathmandu city. The Tibetan name for the site means Sublime Trees'),
(3, current_timestamp, 'Kathmandu Durbar Square', 'Nepal', 1.50, 'Kathmandu Durbar Square (Basantapur Darbar Kshetra) in front of the old royal palace of the former Kathmandu Kingdom is one of three Durbar (royal palace) Squares in the Kathmandu Valley in Nepal, all of which are UNESCO World Heritage Sites. Basantapur is the home of newars.');

ALTER TABLE highlights ADD PRIMARY KEY (id);
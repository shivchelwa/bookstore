CREATE TABLE BOOKS (
                       ISBN char(14) NOT NULL,
                       TITLE varchar(255) NOT NULL,
                       AUTHOR varchar(255) NOT NULL,
                       PRICE decimal(5,2) NOT NULL
);

INSERT INTO BOOKS (ISBN, TITLE, AUTHOR, PRICE) VALUES
                                                   ('978-9386538178', 'As a Man Thinketh', 'James Allen', 16.35),
                                                   ('978-1447277668', 'What I Know for Sure', 'Oprah Winfrey', 25.99),
                                                   ('978-0241491515', 'A Promised Land', 'Barack Obama', 31.99);

ALTER TABLE BOOKS ADD PRIMARY KEY (ISBN);
insert into mst_account_type(account_id, account_name)
values
    ('cash', 'Cash'),
    ('accounts', 'Accounts'),
    ('card', 'Card');
insert into mst_category(category_id, category_name)
values
    ('food', 'Food'),
    ('social_life', 'Social Life'),
    ('pets', 'Pets'),
    ('transport', 'Transport'),
    ('culture', 'Culture'),
    ('hobby', 'Hobby'),
    ('household', 'Household'),
    ('other', 'Other');
insert into mst_transaction_method(method_id, method_name)
values
    ('income', 'Income'),
    ('expense', 'Expense'),
    ('transfer', 'Transfer')
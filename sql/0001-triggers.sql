-- cat sql/0001-triggers.sql | PGPASSWORD=pass psql -h localhost -Umazing financy


DROP TRIGGER IF EXISTS trigger_insert ON charges;
DROP FUNCTION IF EXISTS balance_insert();

CREATE FUNCTION balance_insert() RETURNS trigger AS $$
BEGIN
    IF NEW.kind = 'expense' THEN
        UPDATE account SET total = (total - (NEW.amount / NEW.exchange_rate) ) WHERE account_id = NEW.account_id AND user_id = NEW.user_id;
        UPDATE charges SET balance = (balance - (NEW.amount / NEW.exchange_rate) ) WHERE account_id = NEW.account_id AND user_id = NEW.user_id AND expected_date > NEW.expected_date;
    ELSE
        UPDATE account SET total = (total + (NEW.amount / NEW.exchange_rate) ) WHERE account_id = NEW.account_id AND user_id = NEW.user_id;
        UPDATE charges SET balance = (balance + (NEW.amount / NEW.exchange_rate) ) WHERE account_id = NEW.account_id AND user_id = NEW.user_id AND expected_date > NEW.expected_date;
    END IF;
    NEW.balance := (SELECT total FROM account WHERE account_id = NEW.account_id AND user_id = NEW.user_id);
    RETURN NEW;
END;$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_insert BEFORE INSERT ON charges FOR EACH ROW EXECUTE PROCEDURE balance_insert();



DROP TRIGGER IF EXISTS trigger_delete ON charges;
DROP FUNCTION IF EXISTS balance_delete();

CREATE FUNCTION balance_delete() RETURNS trigger AS $$
BEGIN
    IF OLD.kind = 'expense' THEN
        UPDATE account SET total = (total + (OLD.amount / OLD.exchange_rate) ) WHERE account_id = OLD.account_id AND user_id = OLD.user_id;
        UPDATE charges SET balance = (balance + (OLD.amount / OLD.exchange_rate) ) WHERE account_id = OLD.account_id AND user_id = OLD.user_id AND expected_date > OLD.expected_date;
    ELSE
        UPDATE account SET total = (total - (OLD.amount / OLD.exchange_rate) ) WHERE account_id = OLD.account_id AND user_id = OLD.user_id;
        UPDATE charges SET balance = (balance - (OLD.amount / OLD.exchange_rate) ) WHERE account_id = OLD.account_id AND user_id = OLD.user_id AND expected_date > OLD.expected_date;
    END IF;
    RETURN OLD;
END;$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_delete AFTER DELETE ON charges FOR EACH ROW EXECUTE PROCEDURE balance_delete();


DROP TRIGGER IF EXISTS trigger_update_one ON charges;
DROP FUNCTION IF EXISTS balance_update_one();

CREATE FUNCTION balance_update_one() RETURNS trigger AS $$
BEGIN
    IF NEW.kind = 'expense' THEN
        UPDATE account SET total = (total - (NEW.amount / NEW.exchange_rate) + (OLD.amount / OLD.exchange_rate) ) WHERE account_id = NEW.account_id AND user_id = NEW.user_id;
        UPDATE charges SET balance = (balance - (NEW.amount / NEW.exchange_rate) + (OLD.amount / OLD.exchange_rate) ) WHERE account_id = NEW.account_id AND user_id = NEW.user_id AND expected_date > NEW.expected_date;
    ELSE
        UPDATE account SET total = (total + (NEW.amount / NEW.exchange_rate) - (OLD.amount / OLD.exchange_rate) ) WHERE account_id = NEW.account_id AND user_id = NEW.user_id;
        UPDATE charges SET balance = (balance + (NEW.amount / NEW.exchange_rate) - (OLD.amount / OLD.exchange_rate) ) WHERE account_id = NEW.account_id AND user_id = NEW.user_id AND expected_date > NEW.expected_date;
    END IF;
    NEW.balance := (SELECT total FROM account WHERE account_id = NEW.account_id AND user_id = NEW.user_id);
    RETURN NEW;
END;$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_one BEFORE UPDATE ON charges FOR EACH ROW WHEN (NEW.kind = OLD.kind AND NEW.amount <> OLD.amount) EXECUTE PROCEDURE balance_update_one();


DROP TRIGGER IF EXISTS trigger_update_two ON charges;
DROP FUNCTION IF EXISTS balance_update_two();

CREATE FUNCTION balance_update_two() RETURNS trigger AS $$
BEGIN
    IF NEW.kind = 'expense' THEN
        UPDATE account SET total = (total - (NEW.amount / NEW.exchange_rate) - (OLD.amount / OLD.exchange_rate) ) WHERE account_id = NEW.account_id AND user_id = NEW.user_id;
        UPDATE charges SET balance = (balance - (NEW.amount / NEW.exchange_rate) - (OLD.amount / OLD.exchange_rate) ) WHERE account_id = NEW.account_id AND user_id = NEW.user_id AND expected_date > NEW.expected_date;
    ELSE
        UPDATE account SET total = (total + (NEW.amount / NEW.exchange_rate) + (OLD.amount / OLD.exchange_rate) ) WHERE account_id = NEW.account_id AND user_id = NEW.user_id;
        UPDATE charges SET balance = (balance + (NEW.amount / NEW.exchange_rate) + (OLD.amount / OLD.exchange_rate) ) WHERE account_id = NEW.account_id AND user_id = NEW.user_id AND expected_date > NEW.expected_date;
    END IF;
    NEW.balance := (SELECT total FROM account WHERE account_id = NEW.account_id AND user_id = NEW.user_id);
    RETURN NEW;
END;$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_two BEFORE UPDATE ON charges FOR EACH ROW WHEN (NEW.kind <> OLD.kind) EXECUTE PROCEDURE balance_update_two();

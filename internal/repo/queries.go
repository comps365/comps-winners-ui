package repo

import "fmt"

func buildCompletedLotteriesQuery(prefix string) string {
	return fmt.Sprintf(`
		SELECT p.ID as id, p.post_title AS lottery_name
		FROM %s_posts AS p
			JOIN %s_postmeta pm ON (
				p.ID = pm.post_id 
				AND pm.meta_key  = '_lty_lottery_status'
				AND pm.meta_value = 'lty_lottery_closed'
			)
		WHERE 
			p.post_type = 'product';
	`, prefix, prefix)
}

func buildInstantWinsQuery(prefix string) string {
	return fmt.Sprintf(`
		SELECT pm.meta_value AS instant_win_tickets
		FROM %s_posts AS p
			JOIN %s_postmeta AS pm ON p.ID = pm.post_id
		WHERE
			p.post_parent = ?
			AND p.post_type = 'lty_instant_winners'
			AND pm.meta_key = 'lty_ticket_number'
		GROUP BY instant_win_tickets
	`, prefix, prefix)
}

func buildEntriesQuery(prefix string) string {
	return fmt.Sprintf(`
		SELECT o.billing_email AS email, om.meta_value AS tickets, fst.meta_value as first_name, lst.meta_value as last_name
		FROM %s_wc_orders AS o
			JOIN %s_wc_order_product_lookup opl ON ( o.id = opl.order_id AND opl.product_id = ? ) 
			JOIN %s_wc_orders_meta AS om ON o.id = om.order_id
			JOIN %s_usermeta AS fst ON fst.user_id = o.customer_id
			JOIN %s_usermeta AS lst ON lst.user_id = o.customer_id
		WHERE 
			om.meta_key = 'lty_ticket_ids_in_order'
			AND fst.meta_key = 'first_name'
			AND lst.meta_key = 'last_name'
	`, prefix, prefix, prefix, prefix, prefix)
}

package repo

import "fmt"

func buildCompletedLotteriesQuery(prefix string) string {
	return fmt.Sprintf(`
		SELECT p.ID as id, p.post_title AS lottery_name
		FROM wp_posts AS p
			JOIN %s_postmeta pm ON (
				p.ID = pm.post_id 
				AND pm.meta_key  = '_lty_lottery_status'
				AND pm.meta_value = 'lty_lottery_closed'
			)
		WHERE 
			p.post_type = 'product';
	`, prefix)
}

func buildInstantWinsQuery(prefix string) string {
	return fmt.Sprintf(`
		SELECT pm.meta_value AS instant_win_tickets
		FROM %s_posts AS p
			JOIN wp_postmeta AS pm ON p.ID = pm.post_id
		WHERE
			p.post_parent = ?
			AND p.post_type = 'lty_instant_winners'
			AND pm.meta_key = 'lty_ticket_number'
		GROUP BY instant_win_tickets
	`, prefix)
}

func buildEntriesQuery(prefix string) string {
	return fmt.Sprintf(`
		SELECT o.billing_email AS email, om.meta_value AS entries
		FROM %s_wc_orders AS o
			JOIN wp_wc_order_product_lookup opl ON ( o.id = opl.order_id AND opl.product_id = ?)
			JOIN wp_wc_orders_meta AS om ON o.id = om.order_id
		WHERE 
			om.meta_key = 'lty_ticket_ids_in_order'
	`, prefix)
}

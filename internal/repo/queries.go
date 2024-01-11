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

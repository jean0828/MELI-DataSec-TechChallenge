SELECT 
    CONCAT(customers.first_name, ' ', customers.last_name) AS customer,
    COUNT(events.status) AS failures
FROM 
    customers
JOIN 
    campaigns ON customers.id = campaigns.customer_id
JOIN 
    events ON campaigns.id = events.campaign_id
WHERE 
    events.status = 'failure'
GROUP BY 
    customers.id
HAVING 
    failures > 3
ORDER BY 
    failures DESC;
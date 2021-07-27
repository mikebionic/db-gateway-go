# Database gateway API

API for making direct sql requests to a MSSQL or Postgresql dataabses

> Request for Select query

| url                                   | method |
| ------------------------------------- | :----: |
| 127.0.0.1:8000/api/v1/make-db-request |  POST  |

```json
{
	"query_string": "select * from tbl_mg_materials",
	//"query_string": "select \"ResName\" from tbl_dk_resource",
}
```

**Response**
```json
{
  "data": [
    {
      "T_ID": 1,
      "a_status_id": 1,
      "acc_card_cost_of_sale_id": 108,
      "acc_card_inventory_id": 237,
      "acc_card_purches_disc_id": 0,
      "acc_card_purches_id": 0,
      "acc_card_purches_ret_id": 0,
      "acc_card_sale_disc_id": 36,
      "acc_card_sale_id": 34,
      "acc_card_sale_ret_id": 35,
      "acc_card_scrap_id": 64,
      "acc_card_usage_id": 64,
      "data_send": 0,
      "dept_id": 1,
      "div_id": 1,
      "firm_id": 1,
      "firm_id_guid": "CAB9C297-C332-4BF1-AD3B-A27154BEEC7C",
      "group_code": "SALAT",
      "m_cat_id": 14,
      "mat_auto_price": "0.00000",
      "mat_auto_production": 0,
      "mat_brand_code": "",
      "mat_height": "0.00000",
      "mat_last_purch_arap_id": 0,
      "mat_length": "0.00000",
      "mat_manufacturer": null,
      "mat_name_lang1": "WINIGRET ",
      "mat_name_lang2": "",
      "mat_name_lang3": "",
      "mat_name_lang4": "",
      "mat_name_lang5": "",
      "mat_online_isvisible": null,
      "mat_real_price": "0.00000",
      "mat_shop_code": "",
      "mat_size_code": "",
      "mat_weight": "0.00000",
      "mat_width": "0.00000",
      "material_code": "AN00000049",
      "material_description": "",
      "material_description1": "",
      "material_id": 49,
      "material_id_guid": "d0d43809-d242-43b0-acf7-0e4812bb2e94",
      "material_name": "WINIGRET ",
      "modify_date": "2021-03-18T16:00:33.523Z",
      "security_code": "KUHNYA",
      "spe_code": "",
      "spe_code1": "",
      "spe_code10": "",
      "spe_code11": "",
      "spe_code12": "",
      "spe_code13": null,
      "spe_code14": null,
      "spe_code15": null,
      "spe_code2": "",
      "spe_code3": "",
      "spe_code4": "",
      "spe_code5": "",
      "spe_code6": "",
      "spe_code7": "",
      "spe_code8": "",
      "spe_code9": "",
      "sync_datetime": null,
      "unit_det_id": 1,
      "unit_id": 1
    },
    {
      "T_ID": 1,
      "a_status_id": 1,
      "acc_card_cost_of_sale_id": 108,
      "acc_card_inventory_id": 237,
      "acc_card_purches_disc_id": 0,
      "acc_card_purches_id": 0,
      "acc_card_purches_ret_id": 0,
      "acc_card_sale_disc_id": 36,
      "acc_card_sale_id": 34,
      "acc_card_sale_ret_id": 35,
      "acc_card_scrap_id": 64,
      "acc_card_usage_id": 64,
      "data_send": 0,
      "dept_id": 1,
      "div_id": 1,
      "firm_id": 1,
      "firm_id_guid": "CAB9C297-C332-4BF1-AD3B-A27154BEEC7C",
      "group_code": "SALAT",
      "m_cat_id": 14,
      "mat_auto_price": "0.00000",
      "mat_auto_production": 0,
      "mat_brand_code": "",
      "mat_height": "0.00000",
      "mat_last_purch_arap_id": 0,
      "mat_length": "0.00000",
      "mat_manufacturer": null,
      "mat_name_lang1": "OLWIYE ",
      "mat_name_lang2": "",
      "mat_name_lang3": "",
      "mat_name_lang4": "",
      "mat_name_lang5": "",
      "mat_online_isvisible": null,
      "mat_real_price": "0.00000",
      "mat_shop_code": "",
      "mat_size_code": "",
      "mat_weight": "0.00000",
      "mat_width": "0.00000",
      "material_code": "AN00000050",
      "material_description": "",
      "material_description1": "",
      "material_id": 50,
      "material_id_guid": "CCC697DF-142B-449F-BF40-59C22C02E4CC",
      "material_name": "OLWIYE ",
      "modify_date": "2021-03-18T16:00:33.523Z",
      "security_code": "KUHNYA",
      "spe_code": "",
      "spe_code1": "",
      "spe_code10": "",
      "spe_code11": "",
      "spe_code12": "",
      "spe_code13": null,
      "spe_code14": null,
      "spe_code15": null,
      "spe_code2": "",
      "spe_code3": "",
      "spe_code4": "",
      "spe_code5": "",
      "spe_code6": "",
      "spe_code7": "",
      "spe_code8": "",
      "spe_code9": "",
      "sync_datetime": null,
      "unit_det_id": 1,
      "unit_id": 1
    }
  ],
  "status": 1,
  "total": 2,
  "message": "db query result"
}
```


> Request for Update | Insert | Delete query

| url                                                 | method |
| --------------------------------------------------- | :----: |
| 127.0.0.1:8000/api/v1/make-db-request?executeOnly=1 |  POST  |

```json
{
	"query_string": "update tbl_dk_users set \"URegNo\" = 'SSFK123' where \"UId\" = 1"
}
```
> Response

```json
{
  "data": null,
  "status": 1,
  "total": 1,
  "message": "db query result"
}
```

!! Use **executeOnly** to make other than **SELECT** queries
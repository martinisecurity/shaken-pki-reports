# STIR/SHAKEN CA Ecosystem Compliance

## GBSDTech

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 47 certificates were included in the corpus being tested
- 9 certificates in the corpus were skipped because they are duplicates
- 3 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 35 certificates being tested against the remaining rules
- 1.49 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 932 days is the average remaining validity for the certificates in the corpus
- 949 days is the average initial validity for the certificates in the corpus
- 1 certificates expire in the next 30 days
- 1.00 average number of unexpired certificates per OCN observed
- 35 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_ext_certificate_policies](ISSUES/e_atis_ext_certificate_policies/README.md) | ATIS1000080 |
| 1 | [e_atis_ext_crl_distribution](ISSUES/e_atis_ext_crl_distribution/README.md) | ATIS1000080 |
| 35 | [e_atis_serial_number_size](ISSUES/e_atis_serial_number_size/README.md) | ATIS1000080 |
| 3 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 9 | [e_atis_subject_cn_spc](ISSUES/e_atis_subject_cn_spc/README.md) | ATIS1000080 |
| 1 | [e_atis_tn_auth_list](ISSUES/e_atis_tn_auth_list/README.md) | ATIS1000080 |
| 1 | [e_atis_tn_auth_list_spc_format](ISSUES/e_atis_tn_auth_list_spc_format/README.md) | ATIS1000080 |
| 1 | [e_shaken_certificate_policies_id](ISSUES/e_shaken_certificate_policies_id/README.md) | US_SHAKEN_CP |

#### CA Certificates

- 4 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 4 certificates being tested against the remaining rules
- 2.33 issues on average found in unexpired, trusted, and non-compliant certificates
- 75.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 50.00% of certificates are too old to be assessed against currently enforced expectations
- 6987 days is the average remaining validity for the certificates in the corpus
- 7300 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_ext_authority_key_identifier_ca](ISSUES/e_atis_ext_authority_key_identifier_ca/README.md) | ATIS1000080 |
| 1 | [e_atis_ext_certificate_policies_ca](ISSUES/e_atis_ext_certificate_policies_ca/README.md) | ATIS1000080 |
| 1 | [e_atis_ext_crl_distribution_ca](ISSUES/e_atis_ext_crl_distribution_ca/README.md) | ATIS1000080 |
| 2 | [e_atis_serial_number_size_ca](ISSUES/e_atis_serial_number_size_ca/README.md) | ATIS1000080 |
| 1 | [e_ext_authority_key_identifier_missing](ISSUES/e_ext_authority_key_identifier_missing/README.md) | RFC5280 |
| 1 | [e_ext_authority_key_identifier_no_key_identifier](ISSUES/e_ext_authority_key_identifier_no_key_identifier/README.md) | RFC5280 |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 03&#160;Aug&#160;23&#160;19:55&#160;UTC | MYPBXManager SHAKEN | 02&#160;Aug&#160;26&#160;19:55&#160;UTC | true | [view](CERTS/ea5813855308274fae05fdcae622a159efa47cde2ccf87a9cdf09d9ef43d93f2/README.md) |
| 26&#160;Aug&#160;24&#160;20:14&#160;UTC | allaccesstelecom.com | 25&#160;Aug&#160;25&#160;20:14&#160;UTC | true | [view](CERTS/474b1c6c0e37b3ac6ac6e812a849b99ec02d055f3fad9bad7817e57ad88d150a/README.md) |
| 24&#160;Jan&#160;25&#160;22:34&#160;UTC | Comporium Inc. SHAKEN 0542 | 24&#160;Jan&#160;26&#160;22:34&#160;UTC | true | [view](CERTS/a868bff12eb3b4b7cadc1c66d11cc8aa88521289cdacba0f42257fc627e68515/README.md) |
| 11&#160;Feb&#160;25&#160;18:42&#160;UTC | ccctelecom.com | 10&#160;Feb&#160;26&#160;18:42&#160;UTC | true | [view](CERTS/7ed6c39c5ac9266fe9fcc9b777f7231de3c57af2fa94edc9073ce13a2b61ff99/README.md) |
| 09&#160;Apr&#160;25&#160;22:13&#160;UTC | NovoLink SHAKEN cert | 08&#160;Apr&#160;26&#160;22:13&#160;UTC | true | [view](CERTS/381371aeaf0e18969dae6f24af1c632cea04856e3500a07874f780f928a5e7c0/README.md) |
| 18&#160;Apr&#160;25&#160;03:42&#160;UTC | Allure Telecom Inc. | 17&#160;Apr&#160;26&#160;03:42&#160;UTC | true | [view](CERTS/a618bfb5c97643d797bf1e4833b3137af57f511364cdb41fd990dba3d22b3bd1/README.md) |
| 05&#160;May&#160;25&#160;16:48&#160;UTC | FracTEL LLC SHAKEN | 04&#160;May&#160;28&#160;16:48&#160;UTC | true | [view](CERTS/703a1ebb0715a4bfcacac752f2d1fba1b03aa25e770fbab46df5f89e4d004063/README.md) |
| 06&#160;May&#160;25&#160;21:03&#160;UTC | Eatel SHAKEN 8839 | 05&#160;May&#160;28&#160;21:03&#160;UTC | true | [view](CERTS/45c68a78a546f6dd0dac7228e7b15af43d44da789f78511d8dc9d12b5e3898cf/README.md) |
| 13&#160;May&#160;25&#160;18:15&#160;UTC | Ritter Communications SHAKEN 095A | 12&#160;May&#160;28&#160;18:15&#160;UTC | true | [view](CERTS/6154a0b08f0d0814eb3e0775c0a41f8f805a9a00c7e94852f1b9b0503dae7011/README.md) |
| 13&#160;May&#160;25&#160;19:59&#160;UTC | Lemonweir Valley Telephone Company SHAKEN 0900 | 12&#160;May&#160;28&#160;19:59&#160;UTC | true | [view](CERTS/6a554cb021023b86a78be85a4198cd57c20af3577dc33b73aeed6936f99f7ca8/README.md) |
| 21&#160;May&#160;25&#160;20:21&#160;UTC | Plant Telephone Company SHAKEN 0379 | 20&#160;May&#160;28&#160;20:21&#160;UTC | true | [view](CERTS/b91f977d46aedad90ac0ae9b998fcd59facddd7ae0784322d33dbf587854e931/README.md) |
| 29&#160;May&#160;25&#160;19:05&#160;UTC | Aeneas Communications SHAKEN 2891 | 28&#160;May&#160;28&#160;19:05&#160;UTC | true | [view](CERTS/453946dd067183c43f7f5f64463ce833a9c93291193c488b29ca064279075381/README.md) |
| 30&#160;May&#160;25&#160;12:17&#160;UTC | Priority Communications SHAKEN 327K | 29&#160;May&#160;28&#160;12:17&#160;UTC | true | [view](CERTS/a1634897d10f3dc2ed94719e63be18eb77586c915fa76f8a3cdbfe2b25af17a1/README.md) |
| 03&#160;Jun&#160;25&#160;14:42&#160;UTC | Twin Valley SHAKEN 1840 | 02&#160;Jun&#160;28&#160;14:42&#160;UTC | true | [view](CERTS/ff2f8bbe15cf800b03866277e65b7868d4128a09dfc425faf00923639d5ca2be/README.md) |
| 04&#160;Jun&#160;25&#160;14:14&#160;UTC | Twin Lakes SHAKEN 0579 | 03&#160;Jun&#160;28&#160;14:14&#160;UTC | true | [view](CERTS/28cdaa03feb965d504f83d05650fad329d28c90ffbb15e72418a24b243c24d94/README.md) |
| 09&#160;Jun&#160;25&#160;15:21&#160;UTC | GVTC SHAKEN 2083 | 08&#160;Jun&#160;28&#160;15:21&#160;UTC | true | [view](CERTS/df5698b4fdbc2838ca5b14acbd9f88de9165b55f5c70ae37218178abf00a86ee/README.md) |
| 11&#160;Jun&#160;25&#160;17:11&#160;UTC | INTERBEL TELEPHONE SHAKEN | 10&#160;Jun&#160;26&#160;17:11&#160;UTC | true | [view](CERTS/c5f2080cff46d393050b915ef3c0afd1b42e46b32c128ea5cdc5ec83370d6842/README.md) |
| 12&#160;Jun&#160;25&#160;19:36&#160;UTC | Tohono Oodham Utility Authority SHAKEN 2173 | 11&#160;Jun&#160;28&#160;19:36&#160;UTC | true | [view](CERTS/1e4fb24d1f16ca9e6db4fc0c4978e2c6f4c08a10e1b7d26fc7d9b31bb6991cd5/README.md) |
| 13&#160;Jun&#160;25&#160;20:54&#160;UTC | Segra SHAKEN 1784 | 12&#160;Jun&#160;28&#160;20:54&#160;UTC | true | [view](CERTS/6884bf2d43df25654df4267286d0b64a5236ce6b46ce49d77f299b2d577004ef/README.md) |
| 16&#160;Jun&#160;25&#160;20:44&#160;UTC | Craw Kan Tel Coop SHAKEN 1818 | 15&#160;Jun&#160;28&#160;20:44&#160;UTC | true | [view](CERTS/78fa42ed61e594bc3bb9b590c627f25a4641bdc8bc64de81cd454f0a39f6a3b0/README.md) |
| 17&#160;Jun&#160;25&#160;20:21&#160;UTC | Telesystem SHAKEN 786E | 16&#160;Jun&#160;28&#160;20:21&#160;UTC | true | [view](CERTS/e6caf14534eae46f135535b4fb12da0ba1bfa3484e6189948348110bbd76dd18/README.md) |
| 18&#160;Jun&#160;25&#160;16:27&#160;UTC | Ben Lomand SHAKEN 0553 | 17&#160;Jun&#160;28&#160;16:27&#160;UTC | true | [view](CERTS/b8cc2be99f17446ab5345819a747e05600f0f62517dd3ed0f8fdb83332993886/README.md) |
| 19&#160;Jun&#160;25&#160;19:09&#160;UTC | GCI SHAKEN 7785 | 18&#160;Jun&#160;28&#160;19:09&#160;UTC | true | [view](CERTS/fe444af32c5f43b1f51266e6cda8a8c258eb7260a4f7b4cd6ff89a94d6af0e08/README.md) |
| 20&#160;Jun&#160;25&#160;15:31&#160;UTC | Utility SHAKEN 9262 | 19&#160;Jun&#160;28&#160;15:31&#160;UTC | true | [view](CERTS/bed6418ab21b8318560b44cf2c1336cd7f8e533c48172f5156194e3b33d97bbf/README.md) |
| 20&#160;Jun&#160;25&#160;18:01&#160;UTC | New Horizon SHAKEN 127E | 19&#160;Jun&#160;28&#160;18:01&#160;UTC | true | [view](CERTS/ddeef6001c3bae9328d852daf5ac91c0055229ced11fe0453a143d7953c242fd/README.md) |
| 23&#160;Jun&#160;25&#160;20:37&#160;UTC | FARMERS TELEPHONE 0290 SHAKEN | 22&#160;Jun&#160;26&#160;20:37&#160;UTC | true | [view](CERTS/c8f7b861444acba1f0d4bb972af94ca9656bb37b2e38c2bfe80cf8fc96f5ed34/README.md) |
| 27&#160;Jun&#160;25&#160;12:07&#160;UTC | Valley Telephone Cooperative Inc SHAKEN 2159 | 26&#160;Jun&#160;28&#160;12:07&#160;UTC | true | [view](CERTS/01a9a990d8be7dfdf157de9a2681bf909cf7de6376aa08cee4b1bc3093c2e7c2/README.md) |
| 14&#160;Jul&#160;25&#160;17:00&#160;UTC | Carolina West Wireless SHAKEN 5932 | 13&#160;Jul&#160;28&#160;17:00&#160;UTC | true | [view](CERTS/93b739b8a5d2b8959bbe8dc2b0d9db034a659dbfd8688336dcbf13414d6e3cfe/README.md) |
| 17&#160;Jul&#160;25&#160;19:38&#160;UTC | Crosstel Tandem Inc SEPB SHAKEN 357H | 16&#160;Jul&#160;28&#160;19:38&#160;UTC | true | [view](CERTS/d75c662b84b9758ef4b3662503a6c285c9a272f0c9299c77fcde6ce1e79de832/README.md) |
| 18&#160;Jul&#160;25&#160;14:09&#160;UTC | Point Broadband Inc. Bristol SHAKEN 9809 | 17&#160;Jul&#160;28&#160;14:09&#160;UTC | true | [view](CERTS/ce5307985313ee90816dcde3b6f4a1367e5ea21d42cfd2269acd9597faa64053/README.md) |
| 22&#160;Jul&#160;25&#160;20:53&#160;UTC | Yelcot SHAKEN 1733 | 21&#160;Jul&#160;28&#160;20:53&#160;UTC | true | [view](CERTS/66959aeeb6f4135d90ef5a6c61f081141a74cc1ae170bedfc408a7821c0e0218/README.md) |
| 25&#160;Jul&#160;25&#160;15:43&#160;UTC | Nemont SHAKEN 2247 | 24&#160;Jul&#160;28&#160;15:43&#160;UTC | true | [view](CERTS/4a35a5573560bf76e036e10201f6b8a804255dd39871d9f8eacbc483adedc9db/README.md) |
| 25&#160;Jul&#160;25&#160;20:21&#160;UTC | Everstream SHAKEN 472C | 24&#160;Jul&#160;28&#160;20:21&#160;UTC | true | [view](CERTS/4e96c5a7ab2518957088634d93b4ec9534f08538c517424d02f1590141632ad1/README.md) |
| 30&#160;Jul&#160;25&#160;16:55&#160;UTC | RCN SHAKEN 7615 | 29&#160;Jul&#160;28&#160;16:55&#160;UTC | true | [view](CERTS/48471696816cba0054b8cb47d3a78452fc5721e20a862c5fe25fb88ce37c82f8/README.md) |
| 05&#160;Aug&#160;25&#160;14:11&#160;UTC | Altafiber SHAKEN 600F | 04&#160;Aug&#160;28&#160;14:11&#160;UTC | true | [view](CERTS/f975a59e56f771990398b7717664d96b16e0253a465780f1bd62413f4d6a5f51/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 05&#160;May&#160;21&#160;19:05&#160;UTC | GBSDTech SHAKEN Root CA | 30&#160;Apr&#160;41&#160;19:05&#160;UTC | false | [view](CERTS/6d2bee73a01c1c9fe92273ff8ba56e0c870b7b901cbebcc9e12226fc109e1af9/README.md) |
| 05&#160;May&#160;21&#160;20:22&#160;UTC | 1RouteGroup SHAKEN Intermediate CA | 29&#160;Apr&#160;41&#160;20:22&#160;UTC | true | [view](CERTS/99e9a67644a30ebc33ecc9aa8df6335524d49a4691164e357c5d2406b58a578e/README.md) |
| 27&#160;Dec&#160;23&#160;06:18&#160;UTC | SHAKEN 1RouteGroup Intermediate CA | 22&#160;Dec&#160;43&#160;06:18&#160;UTC | true | [view](CERTS/de6fcb6ba446b8b669a53ed338795ce54b2e3f5be20320fef1644e10fe3b9c3a/README.md) |
| 09&#160;Jul&#160;24&#160;22:39&#160;UTC | SHAKEN 1RouteGroup Intermediate CA | 04&#160;Jul&#160;44&#160;22:39&#160;UTC | true | [view](CERTS/7dea6c98003ca7bd868598f4b62b7be016f5cad5542e8fd95150225d1accf9cf/README.md) |


Generated: 18 Aug 25 21:13 UTC
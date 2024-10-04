# STIR/SHAKEN CA Ecosystem Compliance

## Peeringhub

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 89 certificates were included in the corpus being tested
- 12 certificates in the corpus were skipped because they are duplicates
- 47 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 30 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 3.33% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 261 days is the average remaining validity for the certificates in the corpus
- 260 days is the average initial validity for the certificates in the corpus
- 9 certificates expire in the next 30 days
- 1.03 average number of unexpired certificates per OCN observed
- 29 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_shaken_certificate_policies_id](ISSUES/e_shaken_certificate_policies_id/README.md) | US_SHAKEN_CP |

#### CA Certificates

- 2 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 100.00% of certificates are too old to be assessed against currently enforced expectations
- 5621 days is the average remaining validity for the certificates in the corpus
- 5475 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_subject_cn_ca](ISSUES/e_atis_subject_cn_ca/README.md) | ATIS1000080 |
| 1 | [e_shaken_certificate_policies_id_ca](ISSUES/e_shaken_certificate_policies_id_ca/README.md) | US_SHAKEN_CP |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 24&#160;Oct&#160;23&#160;22:46&#160;UTC | Voipedia SHAKEN 712K | 23&#160;Oct&#160;24&#160;22:46&#160;UTC | true | [view](CERTS/b6029bec3835bab5b57926877ad0ad08198b25153da52b8c7bb40a43ee323b9f/README.md) |
| 05&#160;Jan&#160;24&#160;14:23&#160;UTC | Apex Teleocm LLC_1704464597504 SHAKEN 288K | 29&#160;Nov&#160;24&#160;08:00&#160;UTC | false | [view](CERTS/604ac931870ef2be0b149f5bc4a45c73e6e487f91916b217078efa41d364c34a/README.md) |
| 08&#160;Feb&#160;24&#160;20:24&#160;UTC | Instacall SHAKEN 281K | 22&#160;Jan&#160;25&#160;19:11&#160;UTC | false | [view](CERTS/e280f42a746af003040e83e1e381a1c89a76e7160f3a4649b14543c5e2a1db0d/README.md) |
| 29&#160;Feb&#160;24&#160;19:58&#160;UTC | VaultCom Networks Incorporated_1709236042630 SHAKEN 836K | 28&#160;Feb&#160;25&#160;19:58&#160;UTC | false | [view](CERTS/baad793762bd25655ed76c0f5b50e56da6e18dd26f2a7f0dad52c2606eef574a/README.md) |
| 07&#160;Mar&#160;24&#160;18:12&#160;UTC | Infinity Sip_1709835154271 SHAKEN 279K | 07&#160;Mar&#160;25&#160;18:12&#160;UTC | false | [view](CERTS/36ac9c9983376a0c62d2bd2de4b817debe4798d166442c1e4e72c7356293eceb/README.md) |
| 10&#160;May&#160;24&#160;15:38&#160;UTC | INSTACALL LLC_1715355491325 SHAKEN 281K | 22&#160;Jan&#160;25&#160;19:11&#160;UTC | false | [view](CERTS/ea409ea2f3ce8b166a74e3b0d1a29faa8909b8db7ba7e9c1ded7241f12a87a6b/README.md) |
| 15&#160;May&#160;24&#160;00:00&#160;UTC | VOCALTRANSIT SHAKEN 783J | 15&#160;Oct&#160;24&#160;00:00&#160;UTC | false | [view](CERTS/6f089909297857a1a2de9cb1251c3ac57022775949eef4b610540e8f580466ce/README.md) |
| 28&#160;May&#160;24&#160;19:48&#160;UTC | SHAKEN 088K | 28&#160;May&#160;25&#160;19:48&#160;UTC | false | [view](CERTS/181a97793d7ea7df9a93230c0566a8c6ac2234e21ccc073514161585d8c941cd/README.md) |
| 11&#160;Jun&#160;24&#160;20:32&#160;UTC | Wavecall LLC_1718137942094 SHAKEN 939K | 11&#160;Jun&#160;25&#160;20:32&#160;UTC | false | [view](CERTS/3fac1aaffc1d33baa1e1d342d44d3f4e3ceacaf53a342d56da832f69f5f49168/README.md) |
| 18&#160;Jun&#160;24&#160;21:30&#160;UTC | Ahoi SHAKEN 883K | 18&#160;Jun&#160;25&#160;21:30&#160;UTC | false | [view](CERTS/c041e4169ff485642cd76e764413abb68e951d1ce1e7caea1c187e688363d0cc/README.md) |
| 25&#160;Jun&#160;24&#160;17:14&#160;UTC | Telcast Networks_1719335648799 SHAKEN 611J | 25&#160;Jun&#160;25&#160;17:14&#160;UTC | false | [view](CERTS/b30f6c4c5a9ec72333773acd717df1692bdaf1de3b8e71628e035899149a1d0f/README.md) |
| 06&#160;Jul&#160;24&#160;18:58&#160;UTC | On Air Telecom LLC_1720292281575 SHAKEN 861J | 06&#160;Jul&#160;25&#160;18:58&#160;UTC | false | [view](CERTS/b66429da1f8afc9c2a88652e19e5d693fd965cff8137a189e546b477daab1c9c/README.md) |
| 12&#160;Jul&#160;24&#160;16:44&#160;UTC | VaultTel Solutions Inc_1720802660210 SHAKEN 811K | 25&#160;Jun&#160;25&#160;18:28&#160;UTC | false | [view](CERTS/5eaa8a832b4708f839bdff3d09875030e0536d015f19450d5a628e174defb29e/README.md) |
| 17&#160;Jul&#160;24&#160;16:04&#160;UTC | TruTelco SHAKEN 926K | 20&#160;May&#160;25&#160;17:55&#160;UTC | false | [view](CERTS/552bc06d4306f0f17761764b12fc0955e20640d0875bd950d7eb0f847b99a7cb/README.md) |
| 02&#160;Aug&#160;24&#160;21:33&#160;UTC | Perfect Network LLC_1722634425181 SHAKEN 458K | 02&#160;Aug&#160;25&#160;21:33&#160;UTC | false | [view](CERTS/ab8985f314df3ccccd7986ba9268c047c0059216c830a2d1acf363b8aa5c8859/README.md) |
| 05&#160;Aug&#160;24&#160;03:57&#160;UTC | DiDCentral LLC SHAKEN 756J | 05&#160;Aug&#160;25&#160;03:57&#160;UTC | false | [view](CERTS/69f3edd7be8aa1f21de026ffd129608b918188a75b4bac41bedd9ce94aa92238/README.md) |
| 13&#160;Aug&#160;24&#160;13:01&#160;UTC | Access Tandem SHAKEN 731J | 11&#160;Aug&#160;25&#160;10:14&#160;UTC | false | [view](CERTS/62a1efb1b1c60730d59f83976d022fbea2fc1d1c431f9d3842ff9de04c922a80/README.md) |
| 13&#160;Aug&#160;24&#160;13:32&#160;UTC | Voice Term SHAKEN 240K | 11&#160;Aug&#160;25&#160;09:21&#160;UTC | false | [view](CERTS/c2cbcc1548108e53c296d8f8132613e8884d6b3ef28876df227dc323325c5336/README.md) |
| 19&#160;Aug&#160;24&#160;13:09&#160;UTC | SHAKEN 788J 1724072973 | 19&#160;Aug&#160;25&#160;13:09&#160;UTC | false | [view](CERTS/a93853adcdf9f247570c397d0d5d2af29ba09e0f132ff8c140aa39b326f48b38/README.md) |
| 28&#160;Aug&#160;24&#160;00:27&#160;UTC | DIAL WORLD COMMUNICATIONS LLC_1724804850035 SHAKEN 727K | 05&#160;Aug&#160;25&#160;23:59&#160;UTC | false | [view](CERTS/e035fb9c865457525a27b5e0eccd59de2795a8f3838f366790d7798296786c83/README.md) |
| 04&#160;Sep&#160;24&#160;20:59&#160;UTC | Jaintel LLC SHAKEN 586K | 04&#160;Sep&#160;25&#160;20:59&#160;UTC | false | [view](CERTS/2b291c419ae7f36b4e6626c61b973a82ce6c2f840b90491e142de89571f827a7/README.md) |
| 07&#160;Sep&#160;24&#160;13:46&#160;UTC | TalkAsiaVoip LLC_1725716764206 SHAKEN 198K | 07&#160;Oct&#160;24&#160;13:12&#160;UTC | false | [view](CERTS/297d11cddad0a81790147e6a4227c831c1bf65204aa6b0b1170f70c7343915dc/README.md) |
| 09&#160;Sep&#160;24&#160;02:05&#160;UTC | MAQS Solutions LLC_1725847516773 SHAKEN 255K | 09&#160;Oct&#160;24&#160;20:20&#160;UTC | false | [view](CERTS/b64d4ad2ba78505146f2a26339618f116504d6c140b09c644ba247925b0fb099/README.md) |
| 15&#160;Sep&#160;24&#160;15:39&#160;UTC | HAZTEL INC._1726414788942 SHAKEN 752J | 22&#160;Oct&#160;24&#160;13:56&#160;UTC | false | [view](CERTS/60d86817c4234d39155a84cf2d594a32cd28256bc56237d94c197580516e27c1/README.md) |
| 16&#160;Sep&#160;24&#160;22:47&#160;UTC | INTENEXT TELECOM LLC 1 SHAKEN 650K 1 | 20&#160;Oct&#160;24&#160;18:40&#160;UTC | false | [view](CERTS/9333e82e1e41dd5375a853f36fd3a54a12f0be72064ffb0cc3446f95b6d251ae/README.md) |
| 30&#160;Sep&#160;24&#160;13:11&#160;UTC | ARit services LLC_1727701880593 SHAKEN 827K | 01&#160;Nov&#160;24&#160;15:40&#160;UTC | false | [view](CERTS/93fdb4cbcb3f39df39bba2650f95aac1917477a3a9b24630cf06b76387140b1a/README.md) |
| 30&#160;Sep&#160;24&#160;14:59&#160;UTC | PROCESAMOS LLC SHAKEN 331K. | 31&#160;Oct&#160;24&#160;01:45&#160;UTC | false | [view](CERTS/c12d9c8240fcd038dea3cf2545cc55cba13abd0628a21df54ed4d49a240369aa/README.md) |
| 30&#160;Sep&#160;24&#160;18:34&#160;UTC | Itel Corp SHAKEN 775K 2 | 03&#160;Nov&#160;24&#160;23:15&#160;UTC | false | [view](CERTS/7ade6b21f5f1dce581fb84a30f6a12d3a693309d2a465877d0886de0ccb8b52c/README.md) |
| 30&#160;Sep&#160;24&#160;23:02&#160;UTC | beltalk solutions LLC_1727737340643 SHAKEN 745K | 01&#160;Nov&#160;24&#160;15:02&#160;UTC | false | [view](CERTS/0dacca27bec77e6e8a5a0f2845494a5bf487fdb91fa4bea6936cc95cc6a66d39/README.md) |
| 04&#160;Oct&#160;24&#160;00:00&#160;UTC | Losh Communications, Inc SHAKEN 149K 2024-10-04_000001 | 08&#160;Sep&#160;25&#160;15:03&#160;UTC | false | [view](CERTS/f118b2dc611e66c9223539737104f7d6230a3cee744a87a4ee6ae90ab9e86521/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 17&#160;Dec&#160;20&#160;15:31&#160;UTC | Peeringhub Inc Root CA | 12&#160;Dec&#160;40&#160;15:31&#160;UTC | true | [view](CERTS/0ad4adc0b4d93fdb0e628c577020c73b8a5caff750e7e499f80ee2ab362a3f6a/README.md) |
| 22&#160;Jun&#160;22&#160;22:45&#160;UTC | Peeringhub Inc SHAKEN Intermediate CA 2 | 19&#160;Jun&#160;32&#160;22:45&#160;UTC | true | [view](CERTS/f00871963a40b04269c4b019968e42f9f40964cbfb512ff5342307e9942874ce/README.md) |


Generated: 04 Oct 24 16:29 UTC
# STIR/SHAKEN CA Ecosystem Compliance

## Neustar

### Leaf Certificates

- Average validity span as of leaf certificates 583 days
- Percentage of leaf certificates expiring in the next 30 days is 2.06%
- Certificates with Errors: 82
- Certificates with Warnings: 4
- Certificates with Notices: 0
- Certificates with tests not executed as the requirements were Not Effective at issuance time: 33

| Status | Code | Source | Instances |
|--------|------|--------|-----------|
| not effective | [e_sti_version](ISSUES/e_sti_version/README.md#leaf-certificates) | ATIS-1000080 | 12 |
| not effective | [e_sti_subject_key_identifier](ISSUES/e_sti_subject_key_identifier/README.md#leaf-certificates) | ATIS-1000080 | 12 |
| not effective | [w_cp_1_3_subject_email](ISSUES/w_cp_1_3_subject_email/README.md#leaf-certificates) | United States SHAKEN CP | 32 |
| not effective | [e_sti_authority_key_identifier](ISSUES/e_sti_authority_key_identifier/README.md#leaf-certificates) | ATIS-1000080 | 12 |
| not effective | [e_sti_subject_public_key](ISSUES/e_sti_subject_public_key/README.md#leaf-certificates) | ATIS-1000080 | 12 |
| not effective | [e_sti_tn_auth_list](ISSUES/e_sti_tn_auth_list/README.md#leaf-certificates) | ATIS-1000080 | 12 |
| not effective | [e_sti_crl_distribution](ISSUES/e_sti_crl_distribution/README.md#leaf-certificates) | ATIS-1000080 | 12 |
| error | [e_cp1_3_subject_sn](ISSUES/e_cp1_3_subject_sn/README.md#leaf-certificates) | United States SHAKEN CP | 97 |
| error | [e_sti_certificate_policies](ISSUES/e_sti_certificate_policies/README.md#leaf-certificates) | ATIS-1000080 | 90 |
| not effective | [e_sti_extension_unknown](ISSUES/e_sti_extension_unknown/README.md#leaf-certificates) | ATIS-1000080 | 93 |
| not effective | [e_sti_signature_algorithm](ISSUES/e_sti_signature_algorithm/README.md#leaf-certificates) | ATIS-1000080 | 33 |
| not effective | [e_sti_serial_number](ISSUES/e_sti_serial_number/README.md#leaf-certificates) | ATIS-1000080 | 33 |
| not effective | [e_sti_subject](ISSUES/e_sti_subject/README.md#leaf-certificates) | ATIS-1000080 | 12 |
| not effective | [e_sti_issuer_dn](ISSUES/e_sti_issuer_dn/README.md#leaf-certificates) | ATIS-1000080 | 12 |
| error | [e_sti_crl_distribution_not_reachable](ISSUES/e_sti_crl_distribution_not_reachable/README.md#leaf-certificates) | ATIS-1000080 | 30 |
| not effective | [e_cp1_3_ambiguous_identifier](ISSUES/e_cp1_3_ambiguous_identifier/README.md#leaf-certificates) | United States SHAKEN CP | 34 |
| not effective | [w_cp1_3_subject_rdn_unknown](ISSUES/w_cp1_3_subject_rdn_unknown/README.md#leaf-certificates) | United States SHAKEN CP | 36 |
| not effective | [e_sti_subject_cn](ISSUES/e_sti_subject_cn/README.md#leaf-certificates) | ATIS-1000080 | 34 |
| not effective | [e_sti_key_usage](ISSUES/e_sti_key_usage/README.md#leaf-certificates) | ATIS-1000080 | 12 |
| not effective | [e_sti_basic_constraints](ISSUES/e_sti_basic_constraints/README.md#leaf-certificates) | ATIS-1000080 | 10 |

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\* 1 certificates skipped because they are currently expired.\
\*\*\* 0 certificates skipped because they are not currently trusted by the STI-PA.

#### Issued certificates

| Created at | Name | Problems | Link |
|------------|------|----------|------|
| 18 Dec 19 20:01 UTC | alticeusa.com | false | [view](CERTIFICATES/41c59c1a5cbc1697973cd28b3d6dbb72ce9a2970/README.md) |
| 18 Dec 19 20:08 UTC | vonage | false | [view](CERTIFICATES/9d0d8e4e7eb4b53ea96c66b2a431eb685747f632/README.md) |
| 07 Jan 20 18:48 UTC | peerlessnetwork.com | false | [view](CERTIFICATES/5960121f933226c6968a43965a480d23e36cbc85/README.md) |
| 10 Jan 20 20:09 UTC | CenturyLink.com | false | [view](CERTIFICATES/0a377836df40535e9772f6f0efb3ffb5d9137949/README.md) |
| 29 Jan 20 15:16 UTC | frontier.com | false | [view](CERTIFICATES/d94220a50c87e326070bbd8a4c9d7b7d7f7ab542/README.md) |
| 25 Mar 20 17:29 UTC | cox.com | false | [view](CERTIFICATES/e312112cb4d4e4f7e166acc5cf7bee068d770e13/README.md) |
| 04 Jun 20 18:39 UTC | AGOC | false | [view](CERTIFICATES/c4b3a1125a58b0a1b04eaa08b62324819c90ea9c/README.md) |
| 09 Jul 20 15:29 UTC | digitalipvoice.com | false | [view](CERTIFICATES/8bc020a691adf8cddf209205d987e9e9d7cdf467/README.md) |
| 15 Jul 20 04:13 UTC | SHAKEN | false | [view](CERTIFICATES/49125b9c00a619caeb2345cf7be62aceabfd1b8c/README.md) |
| 20 Aug 20 00:54 UTC | Inteliquent.com | false | [view](CERTIFICATES/9e9abebcc27c7806930008a8b58366be8cd12094/README.md) |
| 15 Sep 20 13:16 UTC | intrado.com | false | [view](CERTIFICATES/eea65b03e54d9f208f4e9e2498857e324d874141/README.md) |
| 25 Sep 20 00:25 UTC | SHAKEN | false | [view](CERTIFICATES/b01a7ccb1bd923ec7a8a10dde0f41249a4643bf5/README.md) |
| 21 Jan 21 18:49 UTC | BigRiverTelephoneCompany | true | [view](CERTIFICATES/3b1c3489397ddc9a5230cf79df2a68616e19f2fe/README.md) |
| 09 Feb 21 21:49 UTC | Ringcentral-ProdKeystore | true | [view](CERTIFICATES/4e63b1c1fe8c8c6bd55b39511d0619ec45f88536/README.md) |
| 16 Feb 21 15:12 UTC | Logix-Keystore | true | [view](CERTIFICATES/2e37e6bc959af5b790ba4727baa232d8a81f6559/README.md) |
| 21 Feb 21 13:57 UTC | Shaken | true | [view](CERTIFICATES/1a059b1bd612bc230679687a7040def8c7358549/README.md) |
| 11 Mar 21 18:18 UTC | SHAKEN-6744 | true | [view](CERTIFICATES/82bd43f274baa78d2403cf4552130195023fdd72/README.md) |
| 12 Mar 21 14:58 UTC | Flowroute | true | [view](CERTIFICATES/920ffa5452d2c9f05ccab954303391049a2ff858/README.md) |
| 24 Mar 21 14:02 UTC | Firstcomm5917 | true | [view](CERTIFICATES/062f8399ad7aa2af6d3f0618d38d32710367b064/README.md) |
| 08 Apr 21 09:47 UTC | Orange | true | [view](CERTIFICATES/ec6c986eac53362c4a2c16b8a19f43cc421fa759/README.md) |
| 16 Apr 21 20:26 UTC | Five9 | true | [view](CERTIFICATES/64603469e1aa730d5eaf318764a32d3b24465568/README.md) |
| 19 Apr 21 23:44 UTC | Warmego | true | [view](CERTIFICATES/744a255bcc98525f4ceb564d2fa06539bba9c2d2/README.md) |
| 03 May 21 21:11 UTC | WindstreamCommunication | true | [view](CERTIFICATES/51cef5d8e5387e2cf7c8cb171737e2cb4216a58c/README.md) |
| 04 May 21 11:46 UTC | intelepeer.com | true | [view](CERTIFICATES/294a4ae29aa6d35e7f3b7a7c84dd65a4ac140b99/README.md) |
| 06 May 21 15:50 UTC | SHAKEN-8468 | true | [view](CERTIFICATES/44e193e2290156b4d91bc3735831fa6afb0ff7ff/README.md) |
| 07 May 21 16:55 UTC | PRD | true | [view](CERTIFICATES/1627721ca177233c1506745d9192ebf730e11fc3/README.md) |
| 18 May 21 15:07 UTC | Granite | true | [view](CERTIFICATES/ea96d15a19dc7ebc2eea690531c91ccc32fb8565/README.md) |
| 19 May 21 14:02 UTC | 846B | true | [view](CERTIFICATES/75927bceda0cdfe72dfa85d3437a170549361bac/README.md) |
| 21 May 21 16:38 UTC | ReInvent | true | [view](CERTIFICATES/918940f013f2d053c5de3c9767df40ad6ff0d3ac/README.md) |
| 01 Nov 21 16:36 UTC | SHAKEN 707J | false | [view](CERTIFICATES/e4369095df0398f1b8736e7c9add7be4917c6892/README.md) |
| 24 Nov 21 19:08 UTC | prod SHAKEN 811J | false | [view](CERTIFICATES/7b0476dd3fc8c512dfe9463bca3d630039030119/README.md) |
| 15 Dec 21 19:36 UTC | SHAKEN 775J | false | [view](CERTIFICATES/d54cf45cb6ba3758042a6a16c739fb9148f93813/README.md) |
| 11 Jan 22 21:31 UTC | SHAKEN 506J | true | [view](CERTIFICATES/ed757497d6919742972670344cc56fde4ff35f24/README.md) |
| 18 Jan 22 18:06 UTC | SHAKEN 899J | true | [view](CERTIFICATES/0c87fe0456cb73c01ca76e3551e55ba7cdba8ddf/README.md) |
| 20 Jan 22 17:34 UTC | SHAKEN 750J | true | [view](CERTIFICATES/dc963db2ed6822a25536cf079c237f08d5ae4fcc/README.md) |
| 02 Feb 22 20:45 UTC | SHAKEN 785J | true | [view](CERTIFICATES/14f50f713574c23c63dce6a2eaf596d5e0f0fc7f/README.md) |
| 04 Mar 22 20:34 UTC | SHAKEN 863J | true | [view](CERTIFICATES/20623088d786b6aae3f3606451a98aaae817b2f6/README.md) |
| 16 Mar 22 13:10 UTC | SHAKEN 393J | true | [view](CERTIFICATES/9bb25f38d97184a5f642a64d54abbdae90e7380f/README.md) |
| 18 Mar 22 18:18 UTC | SHAKEN 804J | true | [view](CERTIFICATES/4730305d08be4a607a60a109487539356f958e92/README.md) |
| 22 Mar 22 14:44 UTC | SHAKEN 701J | true | [view](CERTIFICATES/4bfc404411395872fa4e1bdd75efd22ad56dd5f2/README.md) |
| 22 Mar 22 18:39 UTC | SHAKEN 597F | true | [view](CERTIFICATES/09881d9146c8600778a6d3835c9da6db2b823379/README.md) |
| 24 Mar 22 19:58 UTC | SHAKEN 917J | true | [view](CERTIFICATES/ac5117deaecfcf22b0c90b2f74a932d2879566e9/README.md) |
| 30 Mar 22 00:05 UTC | SHAKEN 5447 | true | [view](CERTIFICATES/3b99807bf67bbd42a1fd9db7d922c808e6ba47c7/README.md) |
| 06 Apr 22 12:06 UTC | SHAKEN 963J | true | [view](CERTIFICATES/af7a793115a18def9ee3e80d6e82afa6df421019/README.md) |
| 08 Apr 22 18:10 UTC | SHAKEN 973J | true | [view](CERTIFICATES/fcb822be75aed3225b93dad5afc7bccae93cfe75/README.md) |
| 11 Apr 22 18:33 UTC | SHAKEN 951J | true | [view](CERTIFICATES/dfc2fd18498333f3836897f2875cb91ecf27e7d2/README.md) |
| 20 Apr 22 20:54 UTC | SHAKEN 966J | true | [view](CERTIFICATES/4d73f8875aba972d8e0378f83c313414b3f4b9a1/README.md) |
| 23 Apr 22 01:57 UTC | SHAKEN 502E | true | [view](CERTIFICATES/e9abb113883eaf7d37812fb3a8b94547e7fa8dcb/README.md) |
| 27 Apr 22 19:37 UTC | SHAKEN 782J | true | [view](CERTIFICATES/9c4232f151c52002be0cec37fcd3ab9cafba26c7/README.md) |
| 02 May 22 13:10 UTC | SHAKEN 854J | true | [view](CERTIFICATES/df67c5ce4f7ebbd1b6b09c7898ea78a9f0630fb2/README.md) |
| 09 May 22 14:53 UTC | SHAKEN 821J | true | [view](CERTIFICATES/8e07c831a4c5e548c6afc8e3551d98aca209c093/README.md) |
| 10 May 22 10:42 UTC | SHAKEN 005K | true | [view](CERTIFICATES/72584ecdb4805a2b263e44c004c3980c8138a0e6/README.md) |
| 10 May 22 18:17 UTC | SHAKEN 402E | true | [view](CERTIFICATES/cd8229035fef92529f5840a90669346e7603bf07/README.md) |
| 13 May 22 14:37 UTC | SHAKEN 473G | true | [view](CERTIFICATES/b03d9d35c96112842650cd868a31a6c974edc206/README.md) |
| 17 May 22 21:17 UTC | SHAKEN 030K | true | [view](CERTIFICATES/480cce3a5e43731ff0f46f7f6f7d84fb8046f72e/README.md) |
| 19 May 22 17:45 UTC | SHAKEN 618J | true | [view](CERTIFICATES/2999c2798ed589378ecde5bfb797406f4babed0c/README.md) |
| 22 May 22 15:34 UTC | SHAKEN 772J | true | [view](CERTIFICATES/d07bd1e44615c07c9b9e110515e7192f31c35981/README.md) |
| 23 May 22 16:43 UTC | SHAKEN 3130 | true | [view](CERTIFICATES/0ca3f33444dfd8387f1b0b3684ae5f95d1d4ca9e/README.md) |
| 24 May 22 15:44 UTC | SHAKEN 845J | true | [view](CERTIFICATES/eb0580b51d673ad99eea75587688d68626e61575/README.md) |
| 24 May 22 16:18 UTC | SHAKEN 4427 | true | [view](CERTIFICATES/b8ace4ec2166c1173851dd7df0ed0cec0493f2a1/README.md) |
| 25 May 22 16:41 UTC | SHAKEN 672J | true | [view](CERTIFICATES/97c75dcbafcb471c6f9700e86efb58f93e5b7d60/README.md) |
| 31 May 22 15:42 UTC | SHAKEN 704J | true | [view](CERTIFICATES/73a4f541c6b374256193ef36dfcd3a46c7ff6b14/README.md) |
| 03 Jun 22 18:24 UTC | ***SHAKEN***464D | true | [view](CERTIFICATES/b9fc91972f86a60e6b515cb176b9616851ffb0ab/README.md) |
| 08 Jun 22 15:19 UTC | SHAKEN 534J | true | [view](CERTIFICATES/759550ad0f6e7086ab7b7a070428d8c7fde19366/README.md) |
| 08 Jun 22 16:50 UTC | SHAKEN 819J | true | [view](CERTIFICATES/1a03eb097db3efef4a7e40f07b53b6cf57ee6954/README.md) |
| 12 Jun 22 15:12 UTC | SHAKEN 261H | true | [view](CERTIFICATES/369619eb89e9326aca1c89cef71b985d30b673de/README.md) |
| 12 Jun 22 15:18 UTC | SHAKEN 033H | true | [view](CERTIFICATES/f13529f6095c87fe53eaa93182d490716370d9ef/README.md) |
| 13 Jun 22 18:51 UTC | SHAKEN 554J | true | [view](CERTIFICATES/ec982d2a606ede7afed8d96a5fbba8bf0334e221/README.md) |
| 16 Jun 22 18:26 UTC | SHAKEN 939H | true | [view](CERTIFICATES/5f432a922d4210b6b32e82cecca79f6cac3606f0/README.md) |
| 17 Jun 22 21:25 UTC | SHAKEN 573J | true | [view](CERTIFICATES/d0f263cc754caa6e444b9f3c350b78fb24f64250/README.md) |
| 21 Jun 22 18:39 UTC | SHAKEN 712J | true | [view](CERTIFICATES/7fe201fe62642765513f645b6ae9e990c77ae56c/README.md) |
| 21 Jun 22 19:56 UTC | SHAKEN 254H | true | [view](CERTIFICATES/ec2b53d5858323c5fa07581e228017c956c006ff/README.md) |
| 21 Jun 22 21:04 UTC | SHAKEN 611J | true | [view](CERTIFICATES/ba2e09f18ce7fcc0b4aea583e77de2819622a21b/README.md) |
| 23 Jun 22 16:36 UTC | SHAKEN 553J | true | [view](CERTIFICATES/4f60323aaec7f2f9e84f17c61bd1aa9bc5bde021/README.md) |
| 23 Jun 22 18:09 UTC | SHAKEN 171K | true | [view](CERTIFICATES/b4a2bb17f5df09697ca5f3ea6ad98001cd9e4b05/README.md) |
| 24 Jun 22 14:05 UTC | SHAKEN 739J | true | [view](CERTIFICATES/df3be214742ba53eecd728a8465a4c60333426da/README.md) |
| 24 Jun 22 19:58 UTC | SHAKEN 067K | true | [view](CERTIFICATES/195517edeeda4afc1a148ef5fe1cdea6e7e38741/README.md) |
| 28 Jun 22 20:15 UTC | SHAKEN 738J | true | [view](CERTIFICATES/c8f1cd8560e9e59ea4f6cb4285552e763b6fcbd2/README.md) |
| 28 Jun 22 21:58 UTC | SHAKEN 743J | true | [view](CERTIFICATES/d80ee0078e6df50ca885cf15a72f1c8a0c7d1c9b/README.md) |
| 29 Jun 22 22:00 UTC | SHAKEN 049K | true | [view](CERTIFICATES/9d5c774d9afe17e1749201abcd597a5760e26d65/README.md) |
| 30 Jun 22 01:51 UTC | SHAKEN 235K | true | [view](CERTIFICATES/b43e0c15e5f51d5acf704d6b29807ba80fcbd11f/README.md) |
| 01 Jul 22 16:21 UTC | SHAKEN 139K | true | [view](CERTIFICATES/181d9981797cf106a0585ca31553e9dab3d9bc00/README.md) |
| 08 Jul 22 17:49 UTC | SHAKEN 709J | true | [view](CERTIFICATES/6d8c5a8630771483ad6ae237a212b38308e838dd/README.md) |
| 13 Jul 22 17:30 UTC | SHAKEN 962J | true | [view](CERTIFICATES/e63044005cf600a61373001653461421f6e465a2/README.md) |
| 13 Jul 22 19:10 UTC | SHAKEN 771J | true | [view](CERTIFICATES/30598eb5437ae2fefdca2a2ecdb6332666d925d2/README.md) |
| 18 Jul 22 17:33 UTC | SHAKEN 715J | true | [view](CERTIFICATES/d0321a392243cdd5e8b2749e4e3883555e2da0cd/README.md) |
| 19 Jul 22 17:18 UTC | SHAKEN 023K | true | [view](CERTIFICATES/d68043eecdd163a62e690d1bbff87cfc72d881ca/README.md) |
| 28 Jul 22 20:56 UTC | SHAKEN 074K | true | [view](CERTIFICATES/8df9e69f0f160ac2c431987be8427329e2039bab/README.md) |
| 08 Aug 22 12:58 UTC | SHAKEN 150K | true | [view](CERTIFICATES/fd83902c0d4196a2464451447f9ccb736378c1ea/README.md) |
| 08 Aug 22 14:30 UTC | SHAKEN 710A | true | [view](CERTIFICATES/5818dcd116e2be68a4cf8186f52e2f46f8c736df/README.md) |
| 18 Aug 22 18:07 UTC | SHAKEN 219K | true | [view](CERTIFICATES/d0133e1ece3d41e56d288b90eb8c9335ec24142e/README.md) |
| 29 Aug 22 19:07 UTC | ATT SHAKEN 4036 | true | [view](CERTIFICATES/baa252858dcd14844f2871eb9b93d60af3692943/README.md) |
| 31 Aug 22 14:47 UTC | SHAKEN 500J | true | [view](CERTIFICATES/eacc4f92800aed1182716331564620b5f6d58106/README.md) |
| 09 Sep 22 14:44 UTC | SHAKEN 5606 | true | [view](CERTIFICATES/e370c45f7048a9986af9a3db5183307c2053c976/README.md) |
| 12 Sep 22 19:52 UTC | SHAKEN 707J | true | [view](CERTIFICATES/01e985113cfde1adfb8122df529c2bae0301fac2/README.md) |
| 15 Sep 22 16:20 UTC | SHAKEN 292K | true | [view](CERTIFICATES/655d6fe195ecd7140c72c94ca361ad4cb893a1cf/README.md) |
| 26 Oct 22 16:36 UTC | SHAKEN 770J | true | [view](CERTIFICATES/ed3545f28b8f2116bd00d59d4635fe94fe65183f/README.md) |

\* For issues relating to this CAs certificate repositories see this [report](URL/README.md).

### CA Certificates

- Certificates with Errors: 3
- Certificates with Warnings: 2
- Certificates with Notices: 10
- Certificates with tests not executed as the requirements were Not Effective at issuance time: 8

| Status | Code | Source | Instances |
|--------|------|--------|-----------|
| not effective | [e_sti_ca_crl_distribution_not_reachable](ISSUES/e_sti_ca_crl_distribution_not_reachable/README.md#ca-certificates) | ATIS-1000080 | 2 |
| not effective | [e_sti_root_extension_unknown](ISSUES/e_sti_root_extension_unknown/README.md#ca-certificates) | ATIS-1000080 | 5 |
| not effective | [e_sti_ca_issuer_dn](ISSUES/e_sti_ca_issuer_dn/README.md#ca-certificates) | ATIS-1000080 | 5 |
| not effective | [e_sti_root_authority_key_identifier](ISSUES/e_sti_root_authority_key_identifier/README.md#ca-certificates) | ATIS-1000080 | 2 |
| not effective | [w_cp1_3_ca_subject_rdn_unknown](ISSUES/w_cp1_3_ca_subject_rdn_unknown/README.md#ca-certificates) | United States SHAKEN CP | 10 |
| not effective | [e_sti_ca_extension_unknown](ISSUES/e_sti_ca_extension_unknown/README.md#ca-certificates) | ATIS-1000080 | 4 |
| not effective | [e_sti_ca_authority_key_identifier](ISSUES/e_sti_ca_authority_key_identifier/README.md#ca-certificates) | ATIS-1000080 | 1 |
| not effective | [e_sti_ca_subject_public_key](ISSUES/e_sti_ca_subject_public_key/README.md#ca-certificates) | ATIS-1000080 | 3 |
| not effective | [e_sti_ca_subject](ISSUES/e_sti_ca_subject/README.md#ca-certificates) | ATIS-1000080 | 3 |
| not effective | [e_sti_ca_serial_number](ISSUES/e_sti_ca_serial_number/README.md#ca-certificates) | ATIS-1000080 | 8 |
| not effective | [e_sti_ca_signature_algorithm](ISSUES/e_sti_ca_signature_algorithm/README.md#ca-certificates) | ATIS-1000080 | 3 |
| notice | [n_pki_ca_key_usage](ISSUES/n_pki_ca_key_usage/README.md#ca-certificates) | SHAKEN PKI Best Practice | 10 |
| not effective | [e_sti_ca_crl_distribution](ISSUES/e_sti_ca_crl_distribution/README.md#ca-certificates) | ATIS-1000080 | 3 |
| not effective | [e_sti_ca_subject_key_identifier](ISSUES/e_sti_ca_subject_key_identifier/README.md#ca-certificates) | ATIS-1000080 | 3 |
| not effective | [e_cp1_3_ca_key_usage_crl_sign](ISSUES/e_cp1_3_ca_key_usage_crl_sign/README.md#ca-certificates) | United States SHAKEN CP | 10 |
| not effective | [e_sti_ca_subject_cn](ISSUES/e_sti_ca_subject_cn/README.md#ca-certificates) | ATIS-1000080 | 8 |
| not effective | [e_sti_basic_constraints](ISSUES/e_sti_basic_constraints/README.md#ca-certificates) | ATIS-1000080 | 3 |
| not effective | [e_sti_ca_version](ISSUES/e_sti_ca_version/README.md#ca-certificates) | ATIS-1000080 | 3 |
| not effective | [e_sti_ca_certificate_policies](ISSUES/e_sti_ca_certificate_policies/README.md#ca-certificates) | ATIS-1000080 | 3 |
| not effective | [n_sti_ca_certificate_policy_critical](ISSUES/n_sti_ca_certificate_policy_critical/README.md#ca-certificates) | ATIS-1000080 | 2 |
| not effective | [e_sti_ca_key_usage](ISSUES/e_sti_ca_key_usage/README.md#ca-certificates) | ATIS-1000080 | 3 |
| not effective | [e_sti_root_certificate_policies](ISSUES/e_sti_root_certificate_policies/README.md#ca-certificates) | ATIS-1000080 | 5 |

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\* 0 certificates skipped because they are currently expired.\
\*\*\* 0 certificates skipped because they are not currently trusted by the STI-PA.

#### Issued certificates

| Created at | Name | Problems | Link |
|------------|------|----------|------|
| 23 Sep 19 13:26 UTC | Neustar Certified Caller ID Root CA | false | [view](CERTIFICATES/028dce43c813a7323688f37a7d491be743d9bbb2/README.md) |
| 23 Sep 19 13:26 UTC | Neustar Certified Caller ID Root CA | false | [view](CERTIFICATES/028dce43c813a7323688f37a7d491be743d9bbb2/README.md) |
| 23 Sep 19 13:32 UTC | Neustar Certified Caller ID CA-1 | false | [view](CERTIFICATES/9c8e8d386d404f4c97d4c37358f2362f01c9cc10/README.md) |
| 08 Oct 20 14:49 UTC | Neustar Canada Certified Caller ID SHAKEN Root CA | false | [view](CERTIFICATES/ba68a7b635a7e85fc1ff99d8f65d7f0ea640dcbd/README.md) |
| 08 Oct 20 14:52 UTC | Neustar Canada Certified Caller ID SHAKEN CA-1 | true | [view](CERTIFICATES/a754e630241fb966de4cac2cd9eb9db90c9d029e/README.md) |
| 17 Aug 21 17:19 UTC | Neustar Certified Caller ID SHAKEN Root CA | false | [view](CERTIFICATES/1eaae3ee5c77b16be8eafe02fb301f376d86a975/README.md) |
| 17 Aug 21 17:19 UTC | Neustar Certified Caller ID SHAKEN Root CA | false | [view](CERTIFICATES/1eaae3ee5c77b16be8eafe02fb301f376d86a975/README.md) |
| 19 Aug 21 03:25 UTC | Neustar Certified Caller ID SHAKEN CA-1 | false | [view](CERTIFICATES/b6f33eebd6fa1f397a22fe4d6300df28960f3061/README.md) |
| 30 Aug 22 06:39 UTC | Neustar Certified Caller ID SHAKEN CA-2 | true | [view](CERTIFICATES/aef2851bc3a7c530c6bcbf864edaeb635cb67155/README.md) |
| 05 Oct 22 17:26 UTC | Neustar Certified Caller ID SHAKEN CA-2 | true | [view](CERTIFICATES/3bbda8c5ef216e5cd2dc6c618a8ebd103aa90e20/README.md) |

Generated: 27/10/2022 at 21:42:52
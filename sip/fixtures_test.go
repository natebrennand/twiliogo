package sip

import (
	"github.com/natebrennand/twiliogo/common"
	"time"
)

var (
	testDomainListFixtureString = `{
	    "first_page_uri": "/2010-04-01/Accounts/ACba8bc05eacf94afdae398e642c9cc32d/SIP/Domains.json?PageSize=50&Page=0",
	    "end": 1,
	    "start": 0,
	    "previous_page_uri": null,
	    "uri": "/2010-04-01/Accounts/ACba8bc05eacf94afdae398e642c9cc32d/SIP/Domains.json?PageSize=50&Page=0",
	    "page_size": 50,
	    "num_pages": 1,
	    "sip_domains": [
	        {
	            "auth_type": "",
	            "voice_status_callback_method": "POST",
	            "subresource_uris": {
	                "ip_access_control_list_mappings": "/2010-04-01/Accounts/ACba8bc05eacf94afdae398e642c9cc32d/SIP/Domains/SD098e7b11c00d0ba152b1d3f084c4b776/IpAccessControlListMappings.json",
	                "credential_list_mappings": "/2010-04-01/Accounts/ACba8bc05eacf94afdae398e642c9cc32d/SIP/Domains/SD098e7b11c00d0ba152b1d3f084c4b776/CredentialListMappings.json"
	            },
	            "date_updated": "Fri, 06 Sep 2013 18:48:50 -0000",
	            "voice_status_callback_url": null,
	            "uri": "/2010-04-01/Accounts/ACba8bc05eacf94afdae398e642c9cc32d/SIP/Domains/SD098e7b11c00d0ba152b1d3f084c4b776.json",
	            "friendly_name": "Scranton Office",
	            "domain_name": "dunder-mifflin-scranton.sip.twilio.com",
	            "account_sid": "ACba8bc05eacf94afdae398e642c9cc32d",
	            "voice_url": "https://dundermifflin.example.com/twilio/app.php",
	            "voice_method": "POST",
	            "sid": "SD098e7b11c00d0ba152b1d3f084c4b776",
	            "date_created": "Fri, 06 Sep 2013 18:48:50 -0000",
	            "voice_fallback_method": "POST",
	            "api_version": "2010-04-01",
	            "voice_fallback_url": null
	        },
	        ...
	    ],
	    "next_page_uri": null,
	    "last_page_uri": "/2010-04-01/Accounts/ACba8bc05eacf94afdae398e642c9cc32d/SIP/Domains.json?PageSize=50&Page=0",
	    "total": 2,
	    "page": 0
	}`

	testDomainListFixture = DomainList{
		ListResponseCore: common.ListResponseCore{
			Start:        0,
			Total:        2,
			NumPages:     1,
			Page:         0,
			PageSize:     50,
			End:          1,
			URI:          "/2010-04-01/Accounts/ACba8bc05eacf94afdae398e642c9cc32d/SIP/Domains.json?PageSize=50&Page=0",
			FirstPageURI: "/2010-04-01/Accounts/ACba8bc05eacf94afdae398e642c9cc32d/SIP/Domains.json?PageSize=50&Page=0",
			LastPageURI:  "/2010-04-01/Accounts/ACba8bc05eacf94afdae398e642c9cc32d/SIP/Domains.json?PageSize=50&Page=0",
		},
		SipDomains: &[]Domain{
			Domain{
				ResourceInfo: common.ResourceInfo{
					Sid:         "SD098e7b11c00d0ba152b1d3f084c4b776",
					AccountSid:  "ACba8bc05eacf94afdae398e642c9cc32d",
					DateCreated: common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
					DateUpdated: common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
					URI:         "/2010-04-01/Accounts/ACba8bc05eacf94afdae398e642c9cc32d/SIP/Domains/SD098e7b11c00d0ba152b1d3f084c4b776.json",
				},
				FriendlyName:              "Scranton Office",
				APIVersion:                "2010-04-01",
				DomainName:                "dunder-mifflin-scranton.sip.twilio.com",
				AuthType:                  "",
				VoiceURL:                  "https://dundermifflin.example.com/twilio/app.php",
				VoiceMethod:               "POST",
				VoiceFallbackMethod:       "POST",
				VoiceStatusCallbackMethod: "POST",
				SubresourceURIs: subresourceURIs{
					IPAccessControlListMappings: "/2010-04-01/Accounts/ACba8bc05eacf94afdae398e642c9cc32d/SIP/Domains/SD098e7b11c00d0ba152b1d3f084c4b776/IpAccessControlListMappings.json",
					CredentialListMappings:      "/2010-04-01/Accounts/ACba8bc05eacf94afdae398e642c9cc32d/SIP/Domains/SD098e7b11c00d0ba152b1d3f084c4b776/CredentialListMappings.json",
				},
			},
		},
	}

	testDomainFixtureString = `{
	    "auth_type": "IP_ACL",
	    "voice_status_callback_method": "POST",
	    "subresource_uris": {
	        "ip_access_control_list_mappings": "/2010-04-01/Accounts/ACba8bc05eacf94afdae398e642c9cc32d/SIP/Domains/SD27f0288630a668bdfbf177f8e22f5ccc/IpAccessControlListMappings.json",
	        "credential_list_mappings": "/2010-04-01/Accounts/ACba8bc05eacf94afdae398e642c9cc32d/SIP/Domains/SD27f0288630a668bdfbf177f8e22f5ccc/CredentialListMappings.json"
	    },
	    "date_updated": "Fri, 06 Sep 2013 19:18:30 -0000",
	    "voice_status_callback_url": null,
	    "uri": "/2010-04-01/Accounts/ACba8bc05eacf94afdae398e642c9cc32d/SIP/Domains/SD27f0288630a668bdfbf177f8e22f5ccc.json",
	    "friendly_name": "Scranton Office",
	    "domain_name": "dunder-mifflin-scranton.sip.twilio.com",
	    "account_sid": "ACba8bc05eacf94afdae398e642c9cc32d",
	    "voice_url": "https://dundermifflin.example.com/twilio/app.php",
	    "voice_method": "POST",
	    "sid": "SD27f0288630a668bdfbf177f8e22f5ccc",
	    "date_created": "Fri, 06 Sep 2013 19:18:30 -0000",
	    "voice_fallback_method": "POST",
	    "api_version": "2010-04-01",
	    "voice_fallback_url": null
	}`

	testDomainFixture = Domain{
		ResourceInfo: common.ResourceInfo{
			Sid:         "SD27f0288630a668bdfbf177f8e22f5ccc",
			AccountSid:  "ACba8bc05eacf94afdae398e642c9cc32d",
			DateCreated: common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
			DateUpdated: common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
			URI:         "/2010-04-01/Accounts/ACba8bc05eacf94afdae398e642c9cc32d/SIP/Domains/SD098e7b11c00d0ba152b1d3f084c4b776.json",
		},
		FriendlyName:              "Scranton Office",
		APIVersion:                "2010-04-01",
		DomainName:                "dunder-mifflin-scranton.sip.twilio.com",
		AuthType:                  "IP_ACL",
		VoiceURL:                  "https://dundermifflin.example.com/twilio/app.php",
		VoiceMethod:               "POST",
		VoiceFallbackMethod:       "POST",
		VoiceStatusCallbackMethod: "POST",
		SubresourceURIs: subresourceURIs{
			IPAccessControlListMappings: "/2010-04-01/Accounts/ACba8bc05eacf94afdae398e642c9cc32d/SIP/Domains/SD098e7b11c00d0ba152b1d3f084c4b776/IpAccessControlListMappings.json",
			CredentialListMappings:      "/2010-04-01/Accounts/ACba8bc05eacf94afdae398e642c9cc32d/SIP/Domains/SD098e7b11c00d0ba152b1d3f084c4b776/CredentialListMappings.json",
		},
	}

	testMappingFixtureString = `{
    	"subresource_uris": {
        	"addresses": "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/IpAccessControlLists/AL95a47094615fe05b7c17e62a7877836c/Addresses.json"
    	},
    	"date_updated": "Wed, 11 Sep 2013 04:06:07 -0000",
    	"friendly_name": "Production Gateways IP Address - Scranton",
    	"uri": "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/IpAccessControlLists/AL95a47094615fe05b7c17e62a7877836c.json",
    	"sid": "AL95a47094615fe05b7c17e62a7877836c",
    	"date_created": "Wed, 11 Sep 2013 04:06:07 -0000"
	}`

	testMappingFixture = Mapping{
		ResourceInfo: common.ResourceInfo{
			Sid:         "AL95a47094615fe05b7c17e62a7877836c",
			DateCreated: common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
			DateUpdated: common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
			URI:         "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/IpAccessControlLists/AL95a47094615fe05b7c17e62a7877836c.json",
		},
		FriendlyName: "Production Gateways IP Address - Scranton",
		SubresourceURIs: addressesSubresourceURI{
			Addresses: "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/IpAccessControlLists/AL95a47094615fe05b7c17e62a7877836c/Addresses.json",
		},
	}

	testCredentialListFixtureString = `{
	    "first_page_uri": "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/Domains/SD32a3c49700934481addd5ce1659f04d2/CredentialListMappings.json?PageSize=50&Page=0",
	    "total": 1,
	    "end": 0,
	    "previous_page_uri": null,
	    "uri": "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/Domains/SD32a3c49700934481addd5ce1659f04d2/CredentialListMappings.json?PageSize=50&Page=0",
	    "page_size": 50,
	    "num_pages": 1,
	    "start": 0,
	    "next_page_uri": null,
	    "credential_list_mappings": [
	        {
	            "subresource_uris": {
	                "credentials": "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/CredentialLists/CL32a3c49700934481addd5ce1659f04d2/Credentials.json"
	            },
	            "date_updated": "Wed, 11 Sep 2013 17:51:38 -0000",
	            "friendly_name": "Production Gateways IP Address - Scranton",
	            "uri": "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/CredentialLists/CL32a3c49700934481addd5ce1659f04d2.json",
	            "account_sid": "AC32a3c49700934481addd5ce1659f04d2",
	            "sid": "CL32a3c49700934481addd5ce1659f04d2",
	            "date_created": "Wed, 11 Sep 2013 17:51:38 -0000"
	        },
	    ],
	    "last_page_uri": "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/Domains/SD32a3c49700934481addd5ce1659f04d2/CredentialListMappings.json?PageSize=50&Page=0",
	    "page": 0
	}`

	testCredentialListFixture = CredentialList{
		ListResponseCore: common.ListResponseCore{
			Start:        0,
			Total:        1,
			NumPages:     1,
			Page:         0,
			PageSize:     50,
			End:          0,
			URI:          "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/Domains/SD32a3c49700934481addd5ce1659f04d2/CredentialListMappings.json?PageSize=50&Page=0",
			FirstPageURI: "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/Domains/SD32a3c49700934481addd5ce1659f04d2/CredentialListMappings.json?PageSize=50&Page=0",
			LastPageURI:  "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/Domains/SD32a3c49700934481addd5ce1659f04d2/CredentialListMappings.json?PageSize=50&Page=0",
		},
		CredentialListMappings: &[]Credential{
			Credential{
				ResourceInfo: common.ResourceInfo{
					Sid:         "CL32a3c49700934481addd5ce1659f04d2",
					AccountSid:  "AC32a3c49700934481addd5ce1659f04d2",
					DateCreated: common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
					DateUpdated: common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
					URI:         "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/CredentialLists/CL32a3c49700934481addd5ce1659f04d2.json",
				},
				FriendlyName: "Production Gateways IP Address - Scranton",
				SubresourceURIs: credentialSubresourceURI{
					Credentials: "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/CredentialLists/CL32a3c49700934481addd5ce1659f04d2/Credentials.json",
				},
			},
		},
	}

	testCredentialFixtureString = `{
	    "subresource_uris": {
	        "credentials": "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/CredentialLists/CL32a3c49700934481addd5ce1659f04d2/Credentials.json"
	    },
	    "date_updated": "Wed, 11 Sep 2013 17:51:38 -0000",
	    "friendly_name": "Production Gateways IP Address - Scranton",
	    "uri": "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/CredentialLists/CL32a3c49700934481addd5ce1659f04d2.json",
	    "account_sid": "AC32a3c49700934481addd5ce1659f04d2",
	    "sid": "CL32a3c49700934481addd5ce1659f04d2",
	    "date_created": "Wed, 11 Sep 2013 17:51:38 -0000"
	}`

	testCredentialFixture = Credential{
		ResourceInfo: common.ResourceInfo{
			Sid:         "CL32a3c49700934481addd5ce1659f04d2",
			AccountSid:  "AC32a3c49700934481addd5ce1659f04d2",
			DateCreated: common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
			DateUpdated: common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
			URI:         "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/CredentialLists/CL32a3c49700934481addd5ce1659f04d2.json",
		},
		FriendlyName: "Production Gateways IP Address - Scranton",
		SubresourceURIs: credentialSubresourceURI{
			Credentials: "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/CredentialLists/CL32a3c49700934481addd5ce1659f04d2/Credentials.json",
		},
	}

	testIPAccessControlListsString = `{
		"first_page_uri": "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/IpAccessControlLists.json?PageSize=50&Page=0",
	    "total": 1,
	    "end": 0,
	    "previous_page_uri": null,
	    "uri": "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/SIP/IpAccessControlLists.json?PageSize=50&Page=0",
	    "page_size": 50,
	    "num_pages": 1,
	    "start": 0,
	    "next_page_uri": null,
	    "ip_access_control_lists": [
	        {
	            "subresource_uris": {
	                "credentials": "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/IPAccessControlLists.json"
	            },
	            "date_updated": "Wed, 11 Sep 2013 17:51:38 -0000",
	            "friendly_name": "Production Gateways IP Address - Scranton",
	            "uri": "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/IPAccessControlLists/AL32a3c49700934481addd5ce1659f04d2.json",
	            "account_sid": "AC32a3c49700934481addd5ce1659f04d2",
	            "sid": "AL32a3c49700934481addd5ce1659f04d2",
	            "date_created": "Wed, 11 Sep 2013 17:51:38 -0000"
	        },
	    ],
	    "last_page_uri": "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/IpAccessControlLists.json?PageSize=50&Page=0",
	    "page": 0
	}`

	testIPAccessControlLists = IPAccessControlLists{
		ListResponseCore: common.ListResponseCore{
			Start:        0,
			Total:        1,
			NumPages:     1,
			Page:         0,
			PageSize:     50,
			End:          0,
			URI:          "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/IPAccessControlLists.json?PageSize=50&Page=0",
			FirstPageURI: "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/IPAccessControlLists.json?PageSize=50&Page=0",
			LastPageURI:  "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/IPAccessControlLists.json?PageSize=50&Page=0",
		},
		ControlLists: &[]IPAccessControlList{
			IPAccessControlList{
				ResourceInfo: common.ResourceInfo{
					Sid:         "AL32a3c49700934481addd5ce1659f04d2",
					URI:         "/2010-04-01/Accounts/AC1fcc43cc0b4157cae6b77cdb692b437e/SIP/IpAccessControlLists/AL32a3c49700934481addd5ce1659f04d2.json",
					DateCreated: common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
					DateUpdated: common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
				},
				SubresourceURIs: addressesSubresourceURI{
					Addresses: "/2010-04-01/Accounts/AC1fcc43cc0b4157cae6b77cdb692b437e/SIP/IpAccessControlLists/AL32a3c49700934481addd5ce1659f04d2/Addresses.json",
				},
			},
		},
	}

	testIPAccessControlListString = `{
	    "subresource_uris": {
	        "addresses": "/2010-04-01/Accounts/AC1fcc43cc0b4157cae6b77cdb692b437e/SIP/IpAccessControlLists/AL0c37d0be69a6a6fe1e270c5fa4a2cac7/Addresses.json"
	    }, 
	    "date_updated": "Wed, 11 Sep 2013 04:06:07 -0000", 
	    "friendly_name": "My new acl", 
	    "uri": "/2010-04-01/Accounts/AC1fcc43cc0b4157cae6b77cdb692b437e/SIP/IpAccessControlLists/AL0c37d0be69a6a6fe1e270c5fa4a2cac7.json", 
	    "sid": "AL0c37d0be69a6a6fe1e270c5fa4a2cac7", 
	    "date_created": "Wed, 11 Sep 2013 04:06:07 -0000"
	}`

	testIPAccessControlList = IPAccessControlList{
		ResourceInfo: common.ResourceInfo{
			Sid:         "AL0c37d0be69a6a6fe1e270c5fa4a2cac7",
			URI:         "/2010-04-01/Accounts/AC1fcc43cc0b4157cae6b77cdb692b437e/SIP/IpAccessControlLists/AL0c37d0be69a6a6fe1e270c5fa4a2cac7.json",
			DateCreated: common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
			DateUpdated: common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
		},
		SubresourceURIs: addressesSubresourceURI{
			Addresses: "/2010-04-01/Accounts/AC1fcc43cc0b4157cae6b77cdb692b437e/SIP/IpAccessControlLists/AL0c37d0be69a6a6fe1e270c5fa4a2cac7/Addresses.json",
		},
	}

	testIPAddressResourceString = `{
	    "date_updated": "Wed, 11 Sep 2013 04:32:30 -0000",
	    "friendly_name": "My office IP Address",
	    "uri": "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/IpAccessControlLists/AL0c37d0be69a6a6fe1e270c5fa4a2cac7/Addresses/IP6fbe2e8619a14fabaefaf4fb131c0b9d.json",
	    "sid": "IP6fbe2e8619a14fabaefaf4fb131c0b9d",
	    "date_created": "Wed, 11 Sep 2013 04:32:30 -0000",
	    "ip_address": "55.102.123.124"
	}`

	testIPAddressResource = IPAddressResource{
		ResourceInfo: common.ResourceInfo{
			Sid:         "AL0c37d0be69a6a6fe1e270c5fa4a2cac7",
			URI:         "/2010-04-01/Accounts/AC1fcc43cc0b4157cae6b77cdb692b437e/SIP/IpAccessControlLists/AL0c37d0be69a6a6fe1e270c5fa4a2cac7.json",
			DateCreated: common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
			DateUpdated: common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
		},
		FriendlyName: "My office IP Address",
		IPAddress:    "55.102.123.124",
	}

	testIPAddressListString = `{
	    "first_page_uri": "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/IpAccessControlLists/AL0c37d0be69a6a6fe1e270c5fa4a2cac7/Addresses.json?PageSize=50&Page=0",
	    "end": 0,
	    "previous_page_uri": null,
	    "ip_addresses": [
	        {
	            "date_updated": "Wed, 11 Sep 2013 04:32:30 -0000",
	            "friendly_name": "My office IP Address",
	            "uri": "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/IpAccessControlLists/AL0c37d0be69a6a6fe1e270c5fa4a2cac7/Addresses/IP6fbe2e8619a14fabaefaf4fb131c0b9d.json",
	            "sid": "IP6fbe2e8619a14fabaefaf4fb131c0b9d",
	            "date_created": "Wed, 11 Sep 2013 04:32:30 -0000",
	            "ip_address": "55.102.123.124"
	        },
	        ...
	    ],
	    "uri": "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/IpAccessControlLists/AL0c37d0be69a6a6fe1e270c5fa4a2cac7/Addresses.json?PageSize=50&Page=0",
	    "page_size": 50,
	    "num_pages": 1,
	    "start": 0,
	    "next_page_uri": null,
	    "last_page_uri": "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/IpAccessControlLists/AL0c37d0be69a6a6fe1e270c5fa4a2cac7/Addresses.json?PageSize=50&Page=0",
	    "total": 1,
	    "page": 0
	}`

	testIPAddressList = IPAddressList{
		ListResponseCore: common.ListResponseCore{
			Start:        0,
			Total:        1,
			NumPages:     1,
			Page:         0,
			PageSize:     50,
			End:          0,
			URI:          "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/IpAccessControlLists/AL0c37d0be69a6a6fe1e270c5fa4a2cac7/Addresses.json?PageSize=50&Page=0",
			FirstPageURI: "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/IpAccessControlLists/AL0c37d0be69a6a6fe1e270c5fa4a2cac7/Addresses.json?PageSize=50&Page=0",
			LastPageURI:  "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/IpAccessControlLists/AL0c37d0be69a6a6fe1e270c5fa4a2cac7/Addresses.json?PageSize=50&Page=0",
		},
		IPAddresses: &[]IPAddressResource{
			IPAddressResource{
				ResourceInfo: common.ResourceInfo{
					Sid:         "IP6fbe2e8619a14fabaefaf4fb131c0b9d",
					URI:         "/2010-04-01/Accounts/AC32a3c49700934481addd5ce1659f04d2/SIP/IpAccessControlLists/AL0c37d0be69a6a6fe1e270c5fa4a2cac7/Addresses/IP6fbe2e8619a14fabaefaf4fb131c0b9d.json",
					DateCreated: common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
					DateUpdated: common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
				},
				FriendlyName: "My office IP Address",
				IPAddress:    "55.102.123.124",
			},
		},
	}
)

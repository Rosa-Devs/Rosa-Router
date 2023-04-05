# cmd

        1 = Set online status to HS node (Ready to accept connections)
                Input:{
                        "cmd": "1",
                        "id": "public key"
                }
                Output:{
                        None
                }

        2 = Get user by id (id is pubkey) (to get user tunnel by id)
                Input:{
                        "id": "public key"
                }
                Output:{
                        "tunnel": "ip:port"
                }
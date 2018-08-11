package quic

//go:generate sh -c "./mockgen_private.sh quic mock_stream_internal_test.go github.com/lucas-clemente/quic-go streamI StreamI"
//go:generate sh -c "./mockgen_private.sh quic mock_receive_stream_internal_test.go github.com/lucas-clemente/quic-go receiveStreamI ReceiveStreamI"
//go:generate sh -c "./mockgen_private.sh quic mock_send_stream_internal_test.go github.com/lucas-clemente/quic-go sendStreamI SendStreamI"
//go:generate sh -c "./mockgen_private.sh quic mock_stream_sender_test.go github.com/lucas-clemente/quic-go streamSender StreamSender"
//go:generate sh -c "./mockgen_private.sh quic mock_stream_getter_test.go github.com/lucas-clemente/quic-go streamGetter StreamGetter"
//go:generate sh -c "./mockgen_private.sh quic mock_stream_frame_source_test.go github.com/lucas-clemente/quic-go streamFrameSource StreamFrameSource"
//go:generate sh -c "./mockgen_private.sh quic mock_crypto_stream_test.go github.com/lucas-clemente/quic-go cryptoStream CryptoStream"
//go:generate sh -c "./mockgen_private.sh quic mock_stream_manager_test.go github.com/lucas-clemente/quic-go streamManager StreamManager"
//go:generate sh -c "./mockgen_private.sh quic mock_unpacker_test.go github.com/lucas-clemente/quic-go unpacker Unpacker"
//go:generate sh -c "./mockgen_private.sh quic mock_quic_aead_test.go github.com/lucas-clemente/quic-go quicAEAD QuicAEAD"
//go:generate sh -c "./mockgen_private.sh quic mock_gquic_aead_test.go github.com/lucas-clemente/quic-go gQUICAEAD GQUICAEAD"
//go:generate sh -c "./mockgen_private.sh quic mock_session_runner_test.go github.com/lucas-clemente/quic-go sessionRunner SessionRunner"
//go:generate sh -c "./mockgen_private.sh quic mock_quic_session_test.go github.com/lucas-clemente/quic-go quicSession QuicSession"
//go:generate sh -c "./mockgen_private.sh quic mock_packet_handler_test.go github.com/lucas-clemente/quic-go packetHandler PacketHandler"
//go:generate sh -c "./mockgen_private.sh quic mock_unknown_packet_handler_test.go github.com/lucas-clemente/quic-go unknownPacketHandler UnknownPacketHandler"
//go:generate sh -c "./mockgen_private.sh quic mock_packet_handler_manager_test.go github.com/lucas-clemente/quic-go packetHandlerManager PacketHandlerManager"
//go:generate sh -c "./mockgen_private.sh quic mock_multiplexer_test.go github.com/lucas-clemente/quic-go multiplexer Multiplexer"
//go:generate sh -c "find . -type f -name 'mock_*_test.go' | xargs sed -i '' 's/quic_go.//g'"
//go:generate sh -c "goimports -w mock*_test.go"

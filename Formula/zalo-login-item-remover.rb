class ZaloLoginItemRemover < Formula
    desc "Theo dõi file log khi mở ứng dụng của Zalo để xóa OpenAtLogin của ứng dụng này."
    homepage "https://iamken.work"
    url "https://github.com/thanhken/homebrew/releases/download/v1.0.2/zalo-login-item-remover"
    sha256 "3907a12fe06915723ed65fd8a1bfa77f67317baa31e22d89d8d5aca148965ed2"
  
    def install
        bin.install "zalo-login-item-remover"
    end

    def post_install
        system "#{bin}/zalo-login-item-remover", "--setup"
    end

    

    def uninstall
        system "#{bin}/zalo-login-item-remover", "--uninstall"
        rm_rf bin/"zalo-login-item-remover"
    end
end
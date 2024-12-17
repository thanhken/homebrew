class ZaloLoginItemRemover < Formula
    desc "Theo dõi file log khi mở ứng dụng của Zalo để xóa OpenAtLogin của ứng dụng này."
    homepage "https://iamken.work"
    url "https://github.com/thanhken/homebrew/releases/download/v1.0.1/zalo-login-item-remover"
    sha256 "df49fc397184f7c7c7eab9c039dafb75d749feae4906bba7f822b8c55b4115af"
  
    def install
        bin.install "zalo-login-item-remover"
    end

    def post_install
        system "#{bin}/zalo-login-item-remover", "--setup", "--bin=#{bin}"
    end

    

    def uninstall
        system "#{bin}/zalo-login-item-remover", "--uninstall"

        # Xóa cmd
        rm_rf bin/"zalo-login-item-remover"

        ohai "Dữ liệu của com.zalo.LoginItemRemover đã được dọn dẹp."
    end
end
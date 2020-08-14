# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
# PLEASE REMOVE ALL GENERATED COMMENTS BEFORE SUBMITTING YOUR PULL REQUEST!
class Juejin < Formula
  desc ""
  homepage ""
  url "https://github.com/youngjuning/homebrew-juejin-spider/raw/master/juejin_0.1.0.tar.gz"
  sha256 "2061517d5df891e0f3405dc10fd4fda72869aea0f7f0157bdc08441ffc5a5afb"
  license ""

  # depends_on "cmake" => :build

  def install
    bin.install "juejin"
  end

  test do
    # `test do` will create, run in and delete a temporary directory.
    #
    # This test will fail and we won't accept that! For Homebrew/homebrew-core
    # this will need to be a test that verifies the functionality of the
    # software. Run the test with `brew test juejin`. Options passed
    # to `brew install` such as `--HEAD` also need to be provided to `brew test`.
    #
    # The installed folder is not in the path, so use the entire path to any
    # executables being tested: `system "#{bin}/program", "do", "something"`.
    system "false"
  end
end

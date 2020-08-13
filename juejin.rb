# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
# PLEASE REMOVE ALL GENERATED COMMENTS BEFORE SUBMITTING YOUR PULL REQUEST!
class Juejin < Formula
  desc ""
  homepage ""
  url "https://github.com/youngjuning/homebrew-juejin-spider/raw/master/juejin_0.0.2.tar.gz"
  sha256 "b3f5ae06fa6c0f85e8ffb04e0327d5d52ab16cf500bf72b73758b6d20ef00df4"
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

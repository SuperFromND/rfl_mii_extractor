# RFL Mii Extractor
RFL-Mii-Extractor is a tool I wrote that can automatically extract Mii data from `RFL_DB.dat`, the internal database file used on the [Mii Channel](https://en.wikipedia.org/wiki/Mii) on the Nintendo Wii. It is written in Go and, although very primitive, should be able to extract Miis from both real Wiis and emulators like Dolphin.

I wrote this tool in order to make it easier to transfer Miis from Dolphin to other platforms; namely, the Nintendo 3DS, in combination with a tool such as [this one](https://kazuki-4ys.github.io/web_apps/MiiInfoEditorCTR/) to conver this data to a QR code.

# Usage
Simply place an `RFL_DB.dat` file in the same folder as the executable and run it. The resulting files will be outputted with a `.mii` extension in that same folder. Sadly, they won't be labelled with anything beyond a number, so it might take some guesswork to determine what Mii's what. I hope to address this in the future!

## Building
I haven't checked to see what version of Go RFL-Mii-Extractor requires, but it should be pretty straightforward to build if you've used Go before:
```bash
git clone https://github.com/SuperFromND/rfl-mii-extractor.git
cd rfl-mii-extractor
make
```
The resulting executable will be placed in the `bin` directory.

## Licensing
RFL-Mii-Extractor is [available under the MIT License.](https://raw.githubusercontent.com/SuperFromND/rfl-mii-extractor/main/LICENSE) <3
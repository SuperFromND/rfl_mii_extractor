# RFL Mii Extractor
RFL-Mii-Extractor is a tool I wrote that can automatically extract Mii data from `RFL_DB.dat`, the internal database file used for the [Mii Channel](https://en.wikipedia.org/wiki/Mii) on the Nintendo Wii. It is written in Go and, though primitive, should be able to extract Miis from both real Wiis and emulators like Dolphin.

I wrote this tool in order to make it easier to transfer Miis from Dolphin to other platforms; namely, the Nintendo 3DS, in combination with a tool such as [this one](https://kazuki-4ys.github.io/web_apps/MiiInfoEditorCTR/) to convert this data to a QR code. It was mostly made for myself (hence why it's very bare-bones), but hopefully it should help someone else out there too!

# Usage
Simply place an `RFL_DB.dat` file in the same folder as the executable and run it. The resulting files will be outputted with a `.mii` extension in that same folder. If you're looking to extract Miis from Dolphin, your `RFL_DB.dat` file can be found in `Dolphin Emulator\Wii\shared2\menu\FaceLib`.

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
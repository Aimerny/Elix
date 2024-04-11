package icu.aimerny.elix;

import icu.aimerny.elix.registry.*;
import net.fabricmc.api.ModInitializer;

public class Elix implements ModInitializer {


    public static final String MOD_ID = "elix";

    @Override
    public void onInitialize() {
        ModItem.init();
        ModItemGroup.init();
        ModBlock.init();

        ModBlockEntity.init();
        ModEnchantment.init();
    }
}

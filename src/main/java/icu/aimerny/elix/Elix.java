package icu.aimerny.elix;

import icu.aimerny.elix.registry.ModBlock;
import icu.aimerny.elix.registry.ModBlockEntity;
import icu.aimerny.elix.registry.ModItem;
import icu.aimerny.elix.registry.ModItemGroup;
import net.fabricmc.api.ModInitializer;

public class Elix implements ModInitializer {


    public static final String MOD_ID = "elix";

    @Override
    public void onInitialize() {
        ModItem.init();
        ModItemGroup.init();
        ModBlock.init();

        ModBlockEntity.init();
    }
}
